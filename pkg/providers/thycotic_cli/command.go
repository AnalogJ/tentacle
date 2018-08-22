package thycotic_cli

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in thycotic secret server via cli",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in thycotic secret server via cli",
				Before: func (ctx *cli.Context) error{
					if !ctx.IsSet("secretId"){
						return fmt.Errorf("`secretId` is required argument")
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					err := p.Authenticate()
					if err != nil {
						return err
					}
					queryData := p.CommandProcessFlagsToQueryData(c)

					secret, err := p.Get(queryData)
					return p.CommandPrintCredentials(c, "get", secret, err)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "secretId",
						Usage:   "Specify the thycotic secret id",
					},
					&cli.StringFlag{
						Name:    "fieldName",
						Usage:   "Specify the thycotic secret field name you would like to retrieve",
						Value:   "password",
					},
				},
			},
		},
	}
}
