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

func (p *Provider) Get(queryData map[string]string) (credentials.BaseInterface, error) {

	resp, err := p.client.Get(queryData["secretid"])
	if err != nil {
		return nil, err
	}

	//determine what type of secret this is, and return that credential type.
	return PopulateCredential(queryData, resp), nil
}

func (p *Provider) List(queryData map[string]string) ([]credentials.BaseInterface, error) {

	resp, err := p.client.List(queryData["criteria"])
	if err != nil {
		return nil, err
	}


	return PopulateSummaryList(queryData, resp), nil
}


func PopulateSummaryList(queryData map[string]string, result api.SearchSecretsResponse) []credentials.BaseInterface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.


	secrets := []credentials.BaseInterface{}

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

func PopulateCredential(queryData map[string]string, result api.GetSecretResponse) credentials.BaseInterface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.

	metadata := map[string]string{}

	//lets start by populating some standard metadata
	metadata["active"] = strconv.FormatBool(result.GetSecretResult.Secret.Active)
	metadata["folderid"] = strconv.Itoa(result.GetSecretResult.Secret.FolderId)
	metadata["secrettypeid"] = strconv.Itoa(result.GetSecretResult.Secret.SecretTypeId) //unfortunately this type can mean different things on different servers.


	// its kind of hard to determine what kind of secret this is, so lets just do some simple/naive processing
	genericSecret := new(credentials.Generic)
	genericSecret.Init()

	secretComponentsNumb := len(result.GetSecretResult.Secret.Items)

	secretdata := map[string]string{}

	for _, item := range result.GetSecretResult.Secret.Items {
		if item.IsNotes && len(item.Value) >0 {
			metadata["Notes"] = item.Value
		} else {
			// fieldname is always lowercased.
			secretdata[strings.ToLower(item.FieldName)] = item.Value
		}
	}

	_, hasUsername := secretdata["username"]
	_, hasPassword := secretdata["password"]

	if secretComponentsNumb == 1 {
		textSecret := new(credentials.Text)
		textSecret.Init()
		textSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		textSecret.Name = result.GetSecretResult.Secret.Name
		textSecret.Data = secretdata
		textSecret.Metadata = metadata
		textSecret.SetText(result.GetSecretResult.Secret.Items[0].Value)
		return textSecret
	} else if hasUsername && hasPassword {
		//this is a username and password secret.
		userpassSecret := new(credentials.UserPass)
		userpassSecret.Init()
		userpassSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		userpassSecret.Name = result.GetSecretResult.Secret.Name
		userpassSecret.Data = secretdata
		userpassSecret.Metadata = metadata
		return userpassSecret
	} else {
		//this is an unknown secret type. Generic.
		genericSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		genericSecret.Name = result.GetSecretResult.Secret.Name
		genericSecret.Metadata = metadata
		return genericSecret
	}
}