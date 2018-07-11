package config

import (
	"github.com/spf13/viper"
	"log"
	"os"

	"tentacle/pkg/utils"
	"tentacle/pkg/errors"
	"tentacle/pkg/providers"
)

// When initializing this class the following methods must be called:
// Config.New
// Config.Init
// This is done automatically when created via the Factory.
type configuration struct {
	*viper.Viper
}


//Viper uses the following precedence order. Each item takes precedence over the item below it:
// explicit call to Set
// flag
// env
// config
// key/value store
// default

func (c *configuration) Init() error {
	c.Viper = viper.New()
	//set defaults

	//if you want to load a non-standard location system config file (~/drawbridge.yml), use ReadConfig
	c.SetConfigType("yaml")
	//c.SetConfigName("drawbridge")
	//c.AddConfigPath("$HOME/")

	//CLI options will be added via the `Set()` function
	return nil
}

func (c *configuration) ReadConfig(configFilePath string) error {
	configFilePath, err := utils.ExpandPath(configFilePath)
	if err != nil {
		return err
	}

	if !utils.FileExists(configFilePath) {
		log.Printf("No configuration file found at %v. Skipping", configFilePath)
		return errors.ConfigFileMissingError("The configuration file could not be found.")
	}

	//log.Printf("Loading configuration file: %s", configFilePath)

	config_data, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("Error reading configuration file: %s", err)
		return err
	}

	err = c.MergeConfig(config_data)
	if err != nil {
		return err
	}

	return nil
}


func (c *configuration) GetProviders() []providers.Interface {

	configs := c.GetStringMap("providers")

	providersList := []providers.Interface{}

	for alias, config := range configs {

		provider, err := providers.Create(alias, config)
		if err != nil {
			return nil
		}
		providersList = append(providersList, provider)
	}
	return providersList

}