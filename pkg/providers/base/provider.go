package base

import (
	"gopkg.in/urfave/cli.v2"
)

type Provider struct {
	Alias          string
	ProviderConfig map[string]interface{}
}

func (p *Provider) Authenticate() error {
	return nil
}

func (p *Provider) Get(queryData map[string]string) error {
	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	return nil
}

//utility/helper functions
func  (p *Provider) CommandFlagsToQueryData(c *cli.Context) map[string]string {
	queryData := map[string]string{}
	for _, flagName := range c.FlagNames() {
		queryData[flagName] = c.String(flagName)
	}
	return queryData
}
