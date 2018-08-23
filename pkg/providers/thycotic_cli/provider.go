package thycotic_cli

import (

"fmt"
"os/exec"
"path"

"github.com/analogj/tentacle/pkg/credentials"
"github.com/analogj/tentacle/pkg/errors"
"github.com/analogj/tentacle/pkg/providers/base"
"github.com/analogj/tentacle/pkg/utils"

)

type Provider struct {
	*base.Provider
	CliHome string
	CliJarFilename string
	CliJarPath string
}

func (p *Provider) Init(alias string, config map[string]interface{}) error {
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias


	// ensure that java is available on the path.
	_, lookErr := exec.LookPath("java")
	if lookErr != nil {
		return errors.DependencyMissingError("java is missing")
	}

	//expand/validate the CLI Home.
	if cliHome, ok := p.ProviderConfig["cli_home"]; ok {
		p.CliHome, _ = utils.ExpandPath(cliHome.(string))
	} else {
		if p.DebugMode {
			fmt.Printf("DEBUG: cli_home path not specified for %v\n", p.Alias)
		}
		return errors.ConfigInvalidError(fmt.Sprintf("CLI Home path is required for %v", p.Alias))
	}


	cliFilename, okFile := p.ProviderConfig["cli_filename"]
	if !okFile || len(cliFilename.(string)) == 0 {
		if p.DebugMode {
			fmt.Printf("DEBUG: cli filename not specified or empty. using 'secretserver-jconsole.jar' for %v\n", p.Alias)
		}
		p.CliJarFilename = "secretserver-jconsole.jar"
	} else {
		p.CliJarFilename = cliFilename.(string)
	}


	// set the Jar Path.
	p.CliJarPath = path.Join(p.CliHome, p.CliJarFilename)
	if !utils.FileExists(p.CliJarPath) {
		if p.DebugMode {
			fmt.Printf("DEBUG: cli jar file not found at %v for %v\n", p.CliJarPath, p.Alias)
		}
		return errors.ConfigInvalidError(fmt.Sprintf("Thycotic jar/cli not found for %v", p.Alias))
	}

	return nil
}

func (p *Provider) Authenticate() error {

	// simple test to ensure that the java + jar combination works. Thycotic CLI does not have a method to test authentication works.
	_, err := utils.SimpleCmdExec("java",[]string{"-jar", p.CliJarPath, "-version"}, p.CliHome , []string{}, !p.DebugMode)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) Get(queryData map[string]string) (credentials.Interface, error) {

	var fieldName string
	if fn, ok := p.ProviderConfig["fieldName"]; ok {
		fieldName = fn.(string)
	} else {
		fieldName = "password"
	}

	// secret retieval
	result, err := utils.SimpleCmdExec("java",[]string{"-jar", p.CliJarPath, "-s", queryData["secretId"], fieldName}, p.CliHome , []string{}, !p.DebugMode)
	if err != nil {
		return nil, err
	}

	return PopulateCredential(queryData, result), nil
}


func PopulateCredential(queryData map[string]string, result string) credentials.Interface {
	//TODO: handle non-text credentials.
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.

	secret := new(credentials.Text)
	secret.Init()
	secret.Data = result

	//set metadata
	secret.Metadata = queryData //sets secretId & fieldName

	return secret
}