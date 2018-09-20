// +build darwin

package keychain

import (
	"gopkg.in/urfave/cli.v2"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.Alias,
		Usage:     "Access secrets and passwords stored on macOS keychain",
		Before: func (ctx *cli.Context) error {
			return p.CommandProcessGlobalFlags(ctx)
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret from macOS keychain",
				//Before: func (ctx *cli.Context) error{
				//	if !ctx.IsSet("id"){
				//		return fmt.Errorf("`id` is required argument")
				//	}
				//	return nil
				//},
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
						Name:    "service",
						Aliases: []string{"where"},
						Usage:   "Specify the keychain secret service",
					},
					&cli.StringFlag{
						Name:    "account",
						Usage:   "Specify the keychain secret account",
					},
					&cli.StringFlag{
						Name:    "label",
						Aliases: []string{"name"},
						Usage:   "Specify the keychain secret label",
					},
					&cli.StringFlag{
						Name:    "description",
						Aliases: []string{"kind"},
						Usage:   "Specify the keychain secret description",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in macOS keychain",
				Action: func(c *cli.Context) error {
					err := p.Authenticate()
					if err != nil {
						return err
					}
					queryData := p.CommandProcessFlagsToQueryData(c)

					secrets, err := p.List(queryData)
					return p.CommandPrintCredentials(c, "list", secrets, err)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "service",
						Aliases: []string{"where"},
						Usage:   "Specify the keychain secret service",
					},
					&cli.StringFlag{
						Name:    "account",
						Usage:   "Specify the keychain secret account",
					},
					&cli.StringFlag{
						Name:    "label",
						Aliases: []string{"name"},
						Usage:   "Specify the keychain secret label",
					},
					&cli.StringFlag{
						Name:    "description",
						Aliases: []string{"kind"},
						Usage:   "Specify the keychain secret description",
					},
				},
			},
		},
	}
}
