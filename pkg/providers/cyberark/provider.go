// +build darwin
package cyberark

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
	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	return nil
}