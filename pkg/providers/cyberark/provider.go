package cyberark

import (
	"log"
	"github.com/analogj/tentacle/pkg/providers/cyberark/api"
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
)

type provider struct {
	*base.Provider
}

func New(alias string, config map[string]interface{}) (*provider, error) {
	p := new(provider)
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias
	return p, nil
}

func (p *provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {


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

	return nil, nil
}