package conjur

import (
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
	"github.com/analogj/tentacle/pkg/errors"
)

type Provider struct {
	*base.Provider

	client *conjurapi.Client
}

func (p *Provider) Init(alias string, config map[string]interface{}) error {
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	return p.ValidateRequireAllOf([]string{"login", "api_key", "appliance_url", "account"}, config)
}

func (p *Provider) Authenticate() error {

	config := conjurapi.LoadConfig()
	config.ApplianceURL = p.ProviderConfig["appliance_url"].(string)
	config.Account = p.ProviderConfig["account"].(string)


	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  p.ProviderConfig["login"].(string),
			APIKey: p.ProviderConfig["api_key"].(string),
		},
	)

	p.client = conjur

	return err
}

func (p *Provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {
	variableId, variableIdOk := queryData["variableid"];
	if  !variableIdOk {
		return nil, errors.InvalidArgumentsError("variableid is empty or invalid")
	}

	respBytes, err := p.client.RetrieveSecret(variableId)

	textSecret := new(credentials.Text)
	textSecret.Init()
	textSecret.Id = variableId
	textSecret.Name = variableId
	textSecret.SetText(string(respBytes))

	if err != nil {
		return nil, err
	}


	return nil, nil
}

func (p *Provider) List(queryData map[string]string) ([]credentials.SummaryInterface, error) {
	return nil, nil
}
