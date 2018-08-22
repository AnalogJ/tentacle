package thycotic_rest

import (
	"github.com/analogj/tentacle/pkg/providers/thycotic_rest/api"
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
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

func (p *Provider) Get(queryData map[string]string) (credentials.Interface, error) {

	client := api.Client {
		Domain: p.ProviderConfig["domain"].(string),
		Hosturl: p.ProviderConfig["hosturl"].(string),
		Username: p.ProviderConfig["username"].(string),
		Password: p.ProviderConfig["password"].(string),
	}

	client.Get(queryData["secretId"])

	return nil, nil
}

func (p *Provider) List(queryData map[string]string) ([]credentials.Interface, error) {
	client := api.Client {
		Domain: p.ProviderConfig["domain"].(string),
		Hosturl: p.ProviderConfig["hosturl"].(string),
		Username: p.ProviderConfig["username"].(string),
		Password: p.ProviderConfig["password"].(string),
	}

	client.List(queryData["criteria"])

	return nil, nil
}