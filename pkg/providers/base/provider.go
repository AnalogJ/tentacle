package base

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"tentacle/pkg/utils"
	"tentacle/pkg/errors"
)

type Provider struct {
	Alias          string
	ProviderConfig map[string]interface{}

	//Global Options
	DebugMode	bool
	OutputMode	string
}

func (p *Provider) Authenticate() error {
	return errors.NotImplementedError("Authenticate function not implemented")
}

func (p *Provider) Get(queryData map[string]string) error {
	return errors.NotImplementedError("Get function not implemented")
}

func (p *Provider) List(queryData map[string]string) error {
	return errors.NotImplementedError("List function not implemented")
}

//utility/helper functions

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

func  (p *Provider) CommandProcessGlobalFlags(c *cli.Context) error {


	p.DebugMode = c.Bool("debug")
	p.OutputMode = c.String("output")

	if p.DebugMode {
		fmt.Println("DEBUG: enabled debug mode")
	}

	return nil
}

func (p *Provider) CommandOutputHelper(c *cli.Context, commandType string, credentialData interface{}, credentialError error) error {
	return nil
}