// +build darwin
package keychain

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func  (p *Provider) Command() *cli.Command {
	return &cli.Command {
		Name:      p.alias,
		Usage:     "Access secrets and passwords stored on macOS keychain",
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
					fmt.Println("secret id: ", c.String(c.FlagNames()[0]))

					p.Authenticate()

					queryData := map[string]string{}
					for _, flagName := range c.FlagNames() {
						queryData[flagName] = c.String(flagName)
					}

					return p.Get(queryData)
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
					fmt.Println("secret id: ", c.String(c.FlagNames()[0]))

					p.Authenticate()

					queryData := map[string]string{}
					for _, flagName := range c.FlagNames() {
						queryData[flagName] = c.String(flagName)
					}

					return p.List(queryData)
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
