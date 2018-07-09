package cyberark

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.alias,
		Usage:     "Access secrets and passwords stored in cyberark vault",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in cyberark vault",
				//Before: func (ctx *cli.Context) error{
				//	if !ctx.IsSet("id"){
				//		return fmt.Errorf("`id` is required argument")
				//	}
				//	return nil
				//},
				Action: func(c *cli.Context) error {
					fmt.Println("secret id: ", c.String(c.FlagNames()[0]))

					p.Authenticate()

					queryData := map[string]string{}
					for _, flagName := range c.FlagNames() {
						queryData[flagName] = c.String(flagName)
					}

					return p.Get(queryData)
				},
				Flags: []cli.Flag{

				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in cyberark vault",
				Action: func(c *cli.Context) error {
					fmt.Println("secret id: ", c.String(c.FlagNames()[0]))

					p.Authenticate()

					queryData := map[string]string{}
					for _, flagName := range c.FlagNames() {
						queryData[flagName] = c.String(flagName)
					}

					return p.List(queryData)
				},
				Flags: []cli.Flag{

				},
			},
		},
	}
}
