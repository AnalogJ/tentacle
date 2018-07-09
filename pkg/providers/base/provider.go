package base

import "gopkg.in/urfave/cli.v2"

type Provider struct {
}


//utility/helper functions
func  (p *Provider) CommandFlagsToQueryData(c *cli.Context) map[string]string {
	queryData := map[string]string{}
	for _, flagName := range c.FlagNames() {
		queryData[flagName] = c.String(flagName)
	}
	return queryData
}
