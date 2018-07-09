// +build darwin
package cyberark

import (
	"log"
	"tentacle/pkg/providers/cyberark/api"
)

type Provider struct {
	alias string
	providerConfig map[string]interface{}
}

func (p *Provider) Init(alias string, config map[string]interface{}) error {
	//validate the config and assign it to providerConfig
	p.providerConfig = config
	p.alias = alias
	return nil
}

func (p *Provider) Authenticate() error {
	return nil
}

func (p *Provider) Get(queryData map[string]string) error {


	client, err := api.NewClient(
		api.SetHost(p.providerConfig["host"].(string)),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	ret, err := client.GetPassword().
		AppID(p.providerConfig["appid"].(string)).
		Safe((p.providerConfig["safe"].(string))).
		Object(queryData["id"]).
		Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	if ret.ErrorCode != "" {
		log.Fatal(ret.ErrorCode)
	}

	log.Println(ret.UserName)
	log.Println(ret.Content)

	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	return nil
}