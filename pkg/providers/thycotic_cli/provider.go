package thycotic_cli

import (
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/analogj/tentacle/pkg/utils"
	"fmt"
	"github.com/analogj/tentacle/pkg/errors"
	"os/exec"
	"path"
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
	err := utils.CmdExec("java",[]string{"-jar", p.CliJarPath, "-version"}, p.CliHome , []string{}, "")
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) Get(queryData map[string]string) (credentials.Interface, error) {

	// secret retieval
	err := utils.CmdExec("java",[]string{"-java", p.CliJarPath, "-s", queryData["secretId"], queryData["fieldName"]}, p.CliHome , []string{}, "")
	if err != nil {
		return nil, err
	}

	return nil, nil
}
