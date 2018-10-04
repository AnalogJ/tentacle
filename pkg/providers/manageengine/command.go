package manageengine

import (
	"gopkg.in/urfave/cli.v2"
	"github.com/analogj/tentacle/pkg/errors"
)


func  (p *provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored in CyberArk Conjur",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret in CyberArk Conjur",
				Before: func (ctx *cli.Context) error{
					idErr := p.CommandValidateRequireAllOf([]string{"id", "resourceid"}, ctx)
					pathErr := p.CommandValidateRequireAllOf([]string{"path"}, ctx)

					if idErr == nil || pathErr == nil {
						return nil
					} else {
						return errors.InvalidArgumentsError("either path or accountId and resourceid must be specified")
					}
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
						Usage:   "Specify the Manage Engine account id",
					},
					&cli.StringFlag{
						Name:    "resourceid",
						Usage:   "Specify the Manage Engine resource id",
					},

					&cli.StringFlag{
						Name:    "path",
						Usage:   "Specify the Manage Engine account path '{resource_name}/{account_name}'",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in CybearArk Conjur",
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
					&cli.StringFlag{
						Name:    "resourceid",
						Usage:   "Specify the Manage Engine resource id",
					},
				},
			},
		},
	}
}