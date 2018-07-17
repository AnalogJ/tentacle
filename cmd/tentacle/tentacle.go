package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/urfave/cli.v2"
	"log"
	"github.com/analogj/tentacle/pkg/config"
	"github.com/analogj/tentacle/pkg/errors"
	"github.com/analogj/tentacle/pkg/version"
	"github.com/fatih/color"
	"github.com/analogj/tentacle/pkg/utils"
)

var goos string
var goarch string

func main() {

	config, err := config.Create()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	//we're going to load the config file manually, since we need to validate it.
	err = config.ReadConfig("~/tentacle.yaml")          // Find and read the config file
	if _, ok := err.(errors.ConfigFileMissingError); ok { // Handle errors reading the config file
		//ignore "could not find config file"
	} else if err != nil {
		os.Exit(1)
	}

	//createFlags, err := createFlags(config)
	//if err != nil {
	//	fmt.Printf("FATAL: %+v\n", err)
	//	os.Exit(1)
	//}

	cli.CommandHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}
USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}
CATEGORY:
   {{.Category}}{{end}}{{if .Description}}
DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	cli.AppHelpTemplate = fmt.Sprintf("%s %s", CustomizeHelpTemplate(), cli.AppHelpTemplate)

	app := &cli.App{
		Name:     "tentacle",
		Usage:    "Base retrieval made simple",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},

		//TODO: add global flag for output type "json, table, raw"
		//TODO: add global flag for debug log level "--debug"
		Flags: []cli.Flag {
			&cli.BoolFlag{
				Name: "debug",
				Value: false,
				Usage: "Enable Debug mode, with extra logging",
			},
			&cli.StringFlag{
				Name: "output",
				Value: "table",
				Usage: "Specify output type. Allowed: 'json', 'table', 'raw'",
			},
		},

		Commands: ConfiguredProviderCommands(config),


	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(color.HiRedString("ERROR: %v", err))
	}
}

func ConfiguredProviderCommands(config config.Interface) []*cli.Command {
	providerList := config.GetProviders()

	commands := []*cli.Command{}

	for _, provider := range providerList {
		commands = append(commands, provider.Command())
	}
	return commands
}

func CustomizeHelpTemplate() string {
	tentacle := "github.com/AnalogJ/tentacle"

	var versionInfo string
	if len(goos) > 0 && len(goarch) > 0 {
		versionInfo = fmt.Sprintf("%s.%s-%s", goos, goarch, version.VERSION)
	} else {
		versionInfo = fmt.Sprintf("dev-%s", version.VERSION)
	}

	subtitle := tentacle + utils.LeftPad2Len(versionInfo, " ", 65-len(tentacle))

	return fmt.Sprintf(utils.StripIndent(
	`
		 ____  ____  __ _  ____  __    ___  __    ____
		(_  _)(  __)(  ( \(_  _)/ _\  / __)(  )  (  __)
		  )(   ) _) /    /  )( /    \( (__ / (_/\ ) _)
		 (__) (____)\_)__) (__)\_/\_/ \___)\____/(____)
		%s
	
		`), subtitle)
}