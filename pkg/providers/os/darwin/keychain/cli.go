// +build darwin
package keychain

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func Cli(nameAlias string, providerConfig ProviderConfig) *cli.Command {
	return &cli.Command {
		Name:      "keychain",
		Usage:     "Access secrets and passwords stored on macOS keychain",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve a specific secret from macOS keychain",
				Before: func (ctx *cli.Context) error{
					if !ctx.IsSet("id"){
						return fmt.Errorf("`id` is required argument")
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					fmt.Println("secret id: ", c.String(c.FlagNames()[0]))

					provider := Provider{ providerConfig:providerConfig }
					provider.Authenticate()
					return provider.Get(map[string]string { "id": c.String(c.FlagNames()[0])})
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "id",
						Usage:   "Specify the account name",
						Value: "",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list all available secrets in macOS keychain",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
		},
	}
}
