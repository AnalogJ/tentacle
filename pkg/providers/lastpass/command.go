package lastpass

import (
	"gopkg.in/urfave/cli.v2"
)


func  (p *provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in Lastpass",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in Lastpass",
				Before: func (ctx *cli.Context) error{
					return p.CommandValidateRequireAllOf([]string{"id"}, ctx)
				},
				Action: func(c *cli.Context) error {
					err := p.Authenticate()
					if err != nil{
						return err
					}
					queryData := p.CommandProcessFlagsToQueryData(c)

					secret, err := p.Get(queryData)
					return p.CommandPrintCredentials(c, "get", secret, err)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "id",
						Usage:   "Specify the Lastpass secret id",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in Lastpass",
				Action: func(c *cli.Context) error {
					err := p.Authenticate()
					if err != nil{
						return err
					}
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