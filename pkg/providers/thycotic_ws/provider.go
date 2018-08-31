package thycotic_ws

import (
	"strconv"
	"strings"

	"github.com/analogj/tentacle/pkg/credentials"
"github.com/analogj/tentacle/pkg/providers/base"
"github.com/analogj/tentacle/pkg/providers/thycotic_ws/api"

)

type Provider struct {
	*base.Provider
	client *api.Client
}

func (p *Provider) Init(alias string, config map[string]interface{}) error {
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	p.client = &api.Client {
		Domain: p.ProviderConfig["domain"].(string),
		Server: p.ProviderConfig["server"].(string),
		Token:  p.ProviderConfig["token"].(string),
		//Username: p.ProviderConfig["username"].(string),
		//Password: p.ProviderConfig["password"].(string),
	}

	//TODO: validate the required configuration is present.
	return nil
}

func (p *Provider) Authenticate() error {



	p.client.Test()

	return nil
}

func (p *Provider) Get(queryData map[string]string) (credentials.Interface, error) {

	resp, err := p.client.Get(queryData["secretid"])
	if err != nil {
		return nil, err
	}

	//determine what type of secret this is, and return that credential type.
	return PopulateCredential(queryData, resp), nil
}

func (p *Provider) List(queryData map[string]string) ([]credentials.Interface, error) {

	resp, err := p.client.List(queryData["criteria"])
	if err != nil {
		return nil, err
	}


	return PopulateSummaryList(queryData, resp), nil
}


func PopulateSummaryList(queryData map[string]string, result api.SearchSecretsResponse) []credentials.Interface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.


	secrets := []credentials.Interface{}

	for _, secret := range result.SearchSecretsResult.SecretSummaries {


		base := new(credentials.Base)
		base.Init()
		base.Id = strconv.Itoa(secret.SecretId)
		base.Name = secret.SecretName
		base.Metadata["secrettypeid"] = strconv.Itoa(secret.SecretTypeId)
		base.Metadata["folderid"] = strconv.Itoa(secret.FolderId)

		secrets = append(secrets, base)
	}

	return secrets
}

func PopulateCredential(queryData map[string]string, result api.GetSecretResponse) credentials.Interface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.

	metadata := map[string]string{}

	//lets start by populating some standard metadata
	metadata["active"] = strconv.FormatBool(result.GetSecretResult.Secret.Active)
	metadata["folderid"] = strconv.Itoa(result.GetSecretResult.Secret.FolderId)
	metadata["secrettypeid"] = strconv.Itoa(result.GetSecretResult.Secret.SecretTypeId) //unfortunately this type can mean different things on different servers.


	// its kind of hard to determine what kind of secret this is, so lets just do some simple/naive processing
	var cred string
	var username string

	secretComponents := len(result.GetSecretResult.Secret.Items)
	if secretComponents > 2 {
		for _, item := range result.GetSecretResult.Secret.Items {
			if item.IsNotes && len(item.Value) >0 {
				metadata["Notes"] = item.Value
			} else if item.IsPassword && len(item.Value) > 0 {
				cred = item.Value
			} else if strings.Contains(strings.ToLower(item.FieldName), "user") && len(item.Value) > 0{
				username = item.Value
			}
		}
	} else if secretComponents == 1 {
		cred = result.GetSecretResult.Secret.Items[0].Value
	}

	if len(cred) >0 && len(username) >0 {
		secret := new(credentials.UserPass)
		secret.Init()
		secret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		secret.Name = result.GetSecretResult.Secret.Name
		secret.SetUsername(username)
		secret.SetPassword(cred)
		secret.Metadata = metadata
		return secret
	} else {
		secret := new(credentials.Text)
		secret.Init()
		secret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		secret.Name = result.GetSecretResult.Secret.Name
		secret.SetText(cred)
		secret.Metadata = metadata
		return secret
	}
}