package thycotic_rest

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in thycotic secret server",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in thycotic secret server",
				Before: func (ctx *cli.Context) error{
					if !ctx.IsSet("name"){
						return fmt.Errorf("`name` is required argument")
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					p.Authenticate()
					queryData := p.CommandProcessFlagsToQueryData(c)

					secret, err := p.Get(queryData)
					return p.CommandPrintCredentials(c, "get", secret, err)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "secretId",
						Usage:   "Specify the thycotic secret id",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in thycotic secret server",
				Action: func(c *cli.Context) error {
					p.Authenticate()
					queryData := p.CommandProcessFlagsToQueryData(c)

					secrets, err := p.List(queryData)
					return p.CommandPrintCredentials(c, "list", secrets, err)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "criteria",
						Aliases: []string{"searchTerm"},
						Usage:   "Specify the thycotic search term",
					},
				},
			},
		},
	}
}