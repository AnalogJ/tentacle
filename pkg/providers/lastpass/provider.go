package lastpass

import (
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	lastpassapi "github.com/analogj/go-lastpass"
	"github.com/analogj/tentacle/pkg/errors"
	"github.com/analogj/tentacle/pkg/constants"
)

type provider struct {
	*base.Provider
	Client *lastpassapi.Client
}

func New(alias string, config map[string]interface{}) (*provider, error) {
	p := new(provider)
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	return p, p.ValidateRequireAllOf([]string{"username", "password"}, config)
}

func (p *provider) Capabilities() map[string]bool {
	return map[string]bool{
		constants.CAP_GET: true,
		constants.CAP_LIST: true,
		constants.CAP_GET_BY_ID: true,
		constants.CAP_CRED_USERPASS: true,
	}
}

func (p *provider) Authenticate() error {

	client := &lastpassapi.Client{}
	client.Init()
	if p.HttpClient != nil {
		client.HttpClient = p.HttpClient
	}

	var multiFactor string
	multiFactorVal, multiFactorOk := p.ProviderConfig["multifactor"]

	if multiFactorOk {
		multiFactor = multiFactorVal.(string)
	} else {
		multiFactor = ""
	}

	err := client.Login(p.ProviderConfig["username"].(string), p.ProviderConfig["password"].(string), multiFactor)

	if err != nil {
		return err
	}

	p.Client = client
	return err
}

func (p *provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {
	id, idOk := queryData["id"];
	if  !idOk {
		return nil, errors.InvalidArgumentsError("id is empty or invalid")
	}

	account, err := p.Client.GetAccount(id)
	if err != nil {
		return nil, err
	}

	userPassSecret := new(credentials.UserPass)
	userPassSecret.Init()
	userPassSecret.Id = account.Id
	userPassSecret.Name = account.Name
	userPassSecret.SetUsername(account.Username)
	userPassSecret.SetPassword(account.Password)
	userPassSecret.Metadata["notes"] = account.Notes
	userPassSecret.Metadata["url"] = account.Url
	userPassSecret.Metadata["group"] = account.Group

	return userPassSecret, nil
}

func (p *provider) List(queryData map[string]string) ([]credentials.SummaryInterface, error) {

	accounts, err := p.Client.GetAccounts()
	if err != nil {
		return nil, err
	}

	summarySecrets := []credentials.SummaryInterface{}

	for _, account := range accounts {
		summary := new(credentials.Summary)
		summary.Init()
		summary.Id = account.Id
		summary.Name = account.Name
		summary.Metadata["username"] = account.Username
		summary.Metadata["notes"] = account.Notes
		summary.Metadata["url"] = account.Url
		summary.Metadata["group"] = account.Group

		summarySecrets = append(summarySecrets, summary)

	}

	return summarySecrets, nil
}
