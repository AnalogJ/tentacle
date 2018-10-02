package cyberark

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in cyberark vault",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in cyberark vault",
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
						Name:    "name",
						Aliases: []string{"account"},
						Usage:   "Specify the cyberark secret name",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in cybearark vault",
				Action: func(c *cli.Context) error {
					p.Authenticate()
					queryData := p.CommandProcessFlagsToQueryData(c)

					secrets, err := p.List(queryData)
					return p.CommandPrintCredentials(c, "list", secrets, err)

				},
				Flags: []cli.Flag{

				},
			},
		},
	}
}
