package conjur

import (
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
	"github.com/analogj/tentacle/pkg/errors"
	"github.com/analogj/tentacle/pkg/constants"
)

type provider struct {
	*base.Provider

	client *conjurapi.Client
}

func New(alias string, config map[string]interface{}) (*provider, error) {
	p := new(provider)

	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	return p, p.ValidateRequireAllOf([]string{"login", "api_key", "appliance_url", "account"}, config)
}

func (p *provider) Capabilities() map[string]bool {
	return map[string]bool{
		constants.CAP_GET: true,
		constants.CAP_GET_BY_ID: true,
		constants.CAP_CRED_TEXT: true,
	}
}

func (p *provider) Authenticate() error {

	config, err := conjurapi.LoadConfig()
	if err != nil {
		return err
	}
	config.ApplianceURL = p.ProviderConfig["appliance_url"].(string)
	config.Account = p.ProviderConfig["account"].(string)


	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  p.ProviderConfig["login"].(string),
			APIKey: p.ProviderConfig["api_key"].(string),
		},
	)

	if p.HttpClient != nil {
		conjur.SetHttpClient(p.HttpClient)

	}
	p.client = conjur

	return err
}

func (p *provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {

	variableId, variableIdOk := queryData["id"]
	if  !variableIdOk {
		return nil, errors.InvalidArgumentsError("id is empty or invalid")
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

	return textSecret, nil
}
