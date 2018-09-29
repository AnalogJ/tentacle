package base

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"github.com/analogj/tentacle/pkg/utils"
	"github.com/analogj/tentacle/pkg/errors"
	"github.com/analogj/tentacle/pkg/credentials"
	"encoding/json"
	"strings"
)

type Provider struct {
	Alias          string
	ProviderConfig map[string]interface{}

	//Global Options
	DebugMode	bool
	OutputMode	string
}

func (p *Provider) Authenticate() error {
	return errors.NotImplementedError("Authenticate action is unsupported by this provider.")
}

func (p *Provider) Get(queryData map[string]string) error {
	return errors.NotImplementedError("Get action is unsupported by this provider.")
}

func (p *Provider) List(queryData map[string]string)  ([]credentials.SummaryInterface, error) {
	return nil, errors.NotImplementedError("List action is unsupported by this provider.")
}


//utility/helper functions

func (p *Provider) ValidateRequireOneOf(oneOf []string, data map[string]interface{}) error {

	var oneSet = false
	for _, flag := range oneOf {
		if _, ok := data[flag]; ok {
			oneSet = true
			break
		}
	}

	if !oneSet {
		return errors.InvalidArgumentsError(fmt.Sprintf("One of `%s` is required", strings.Join(oneOf, "` or `")))
	} else {
		return nil
	}
}

func (p *Provider) ValidateRequireAllOf(allOf []string, data map[string]interface{}) error {

	var allSet = true
	for _, flag := range allOf {
		if _, ok := data[flag]; !ok {
			allSet = false
		}
	}

	if !allSet {
		return errors.InvalidArgumentsError(fmt.Sprintf("`%s` are required", strings.Join(allOf, "` and `")))
	} else {
		return nil
	}
}

func (p *Provider) CommandValidateRequireOneOf(oneOf []string, c *cli.Context) error {

	var oneSet = false
	for _, flag := range oneOf {
		if c.IsSet(flag) {
			 oneSet = true
			 break
		}
	}

	if !oneSet {
		return errors.InvalidArgumentsError(fmt.Sprintf("One of `%s` is required", strings.Join(oneOf, "` or `")))
	} else {
		return nil
	}
}

func (p *Provider) CommandValidateRequireAllOf(allOf []string, c *cli.Context) error {

	var allSet = true
	for _, flag := range allOf {
		if !c.IsSet(flag) {
			allSet = false
		}
	}

	if !allSet {
		return errors.InvalidArgumentsError(fmt.Sprintf("`%s` are required", strings.Join(allOf, "` and `")))
	} else {
		return nil
	}
}


var reservedFlags = []string{ "output", "debug" }

func  (p *Provider) CommandProcessFlagsToQueryData(c *cli.Context) map[string]string {


	queryData := map[string]string{}
	for _, flagName := range c.FlagNames() {
		//skip over reserved flags.
		if utils.SliceIncludes(reservedFlags, flagName){
			continue
		}

		queryData[flagName] = c.String(flagName)
	}

	if p.DebugMode{
		fmt.Printf("DEBUG: Query data: %#v\n", queryData)
	}

	return queryData
}

func (p *Provider) CommandProcessGlobalFlags(c *cli.Context) error {


	p.DebugMode = c.Bool("debug")
	p.OutputMode = c.String("output")

	if p.DebugMode {
		fmt.Println("DEBUG: enabled debug mode")
	}

	return nil
}

func (p *Provider) CommandPrintCredentials(c *cli.Context, commandType string, credentialData interface{}, credentialError error) error {

	if credentialError != nil {
		return credentialError
	}
	switch commandType {
	case "get":
		return PrintCredential(p.OutputMode, credentialData.(credentials.GenericInterface))
	case "list":
		return PrintCredentials(p.OutputMode, credentialData.([]credentials.SummaryInterface))
	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("command type is invalid (%v)", commandType))
	}

	return nil
}

func PrintCredential(outputMode string, credential credentials.GenericInterface) error {

	switch outputMode {
	case "json":
		secret, err := credential.ToJsonString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "yaml":
		secret, err := credential.ToYamlString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "raw":
		secret, err := credential.ToRawString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "table":
		secret, err := credential.ToTableString()
		if err != nil {
			return err
		}
		fmt.Println(secret)

	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("output mode is invalid (%v)", outputMode))
	}
	return nil
}

func PrintCredentials(outputMode string, credentials []credentials.SummaryInterface) error {

	switch outputMode {
	case "json":
		jsonBytes, err := json.MarshalIndent(credentials,"", "    ")
		if err != nil {
			return err
		}
		fmt.Printf(string(jsonBytes))
	case "raw":
		return errors.InvalidArgumentsError("output mode raw is unsupported for list")
	case "table":
		//TODO: print a table.
		//fmt.Println(credential.ToTableString())
	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("output mode is invalid (%v)", outputMode))
	}
	return nil
}