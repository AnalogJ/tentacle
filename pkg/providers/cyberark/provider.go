// +build darwin
package cyberark

import (
	"log"
	"tentacle/pkg/providers/cyberark/api"
	"tentacle/pkg/providers/base"
)

type Provider struct {
	*base.Provider
}

func (p *Provider) Init(alias string, config map[string]interface{}) error {
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias
	return nil
}

func (p *Provider) Get(queryData map[string]string) error {


	client, err := api.NewClient(
		api.SetHost(p.ProviderConfig["host"].(string)),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	ret, err := client.GetPassword().
		AppID(p.ProviderConfig["appid"].(string)).
		Safe((p.ProviderConfig["safe"].(string))).
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