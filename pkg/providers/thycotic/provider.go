package thycotic

import (
	"tentacle/pkg/providers/thycotic/api"
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

	client := api.Client {
		Domain: p.ProviderConfig["domain"].(string),
		Hosturl: p.ProviderConfig["hosturl"].(string),
		Username: p.ProviderConfig["username"].(string),
		Password: p.ProviderConfig["password"].(string),
	}

	client.Get(queryData["secretId"])

	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	client := api.Client {
		Domain: p.ProviderConfig["domain"].(string),
		Hosturl: p.ProviderConfig["hosturl"].(string),
		Username: p.ProviderConfig["username"].(string),
		Password: p.ProviderConfig["password"].(string),
	}

	client.List(queryData["criteria"])

	return nil
}