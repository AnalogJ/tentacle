package thycotic

import (
	"tentacle/pkg/providers/thycotic/api"
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

	client := api.Client {
		Domain: p.providerConfig["domain"].(string),
		Hosturl: p.providerConfig["hosturl"].(string),
		Username: p.providerConfig["username"].(string),
		Password: p.providerConfig["password"].(string),
	}

	client.Get(queryData["secretId"])

	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	client := api.Client {
		Domain: p.providerConfig["domain"].(string),
		Hosturl: p.providerConfig["hosturl"].(string),
		Username: p.providerConfig["username"].(string),
		Password: p.providerConfig["password"].(string),
	}

	client.List(queryData["criteria"])

	return nil
}