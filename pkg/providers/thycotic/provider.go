package thycotic

import (
	"strconv"
	"strings"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/providers/thycotic/api"
	"github.com/analogj/tentacle/pkg/constants"
)

type provider struct {
	*base.Provider
	client *api.Client
}

func New(alias string, config map[string]interface{}) (*provider, error) {
	p := new(provider)
	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	verr := p.ValidateRequireAllOf([]string{"domain", "server", "token"}, config)
	if verr != nil {
		return nil, verr
	}

	p.client = new(api.Client)
	p.client.Init(p.ProviderConfig["domain"].(string), p.ProviderConfig["server"].(string), p.ProviderConfig["token"].(string))
	//if p.HttpClient != nil {
	//	p.client.HttpClient = p.HttpClient
	//}
	return p, nil
}

func (p *provider) Capabilities() map[string]bool {
	return map[string]bool{
		constants.CAP_GET: true,
		constants.CAP_LIST: true,
		constants.CAP_GET_BY_ID: true,
		constants.CAP_GET_BY_PATH: true,
		constants.CAP_CRED_USERPASS: true,
		constants.CAP_CRED_SSH: true,
	}
}

func (p *provider) Authenticate() error {
	_, err := p.client.Test()
	return err
}

func (p *provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {

	var resp api.GetSecretResponse
	var err error
	if val, ok := queryData["id"]; ok {
		resp, err = p.client.GetById(val)
		if err != nil {
			return nil, err
		}
	} else {
		resp, err = p.client.GetByPath(queryData["path"])
		if err != nil {
			return nil, err
		}
	}

	//determine what type of secret this is, and return that credential type.
	return p.populateCredential(queryData, resp), nil
}

func (p *provider) List(queryData map[string]string) ([]credentials.SummaryInterface, error) {

	resp, err := p.client.List(queryData["criteria"])
	if err != nil {
		return nil, err
	}


	return p.populateSummaryList(queryData, resp), nil
}


func (p *provider) populateSummaryList(queryData map[string]string, result api.SearchSecretsResponse) []credentials.SummaryInterface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.


	secrets := []credentials.SummaryInterface{}

	for _, secret := range result.SearchSecretsResult.SecretSummaries {


		//base := credentials.CreateSummary()
		base := new(credentials.Summary)
		base.Init()
		base.Id = strconv.Itoa(secret.SecretId)
		base.Name = secret.SecretName
		base.Metadata["secrettypeid"] = strconv.Itoa(secret.SecretTypeId)
		base.Metadata["folderid"] = strconv.Itoa(secret.FolderId)

		secrets = append(secrets, base)
	}

	return secrets
}

func (p *provider) populateCredential(queryData map[string]string, result api.GetSecretResponse) credentials.GenericInterface {
	// As of now, theres no way to determine what type of credential we've recieved, always return a Text type.

	metadata := map[string]string{}

	//lets start by populating some standard metadata
	metadata["active"] = strconv.FormatBool(result.GetSecretResult.Secret.Active)
	metadata["folderid"] = strconv.Itoa(result.GetSecretResult.Secret.FolderId)
	metadata["secrettypeid"] = strconv.Itoa(result.GetSecretResult.Secret.SecretTypeId) //unfortunately this type can mean different things on different servers.


	// its kind of hard to determine what kind of secret this is, so lets just do some simple/naive processing
	secretdata := map[string]string{}

	for _, item := range result.GetSecretResult.Secret.Items {

		if item.IsNotes && len(item.Value) >0 {
			metadata["notes"] = item.Value
		} else if item.IsFile {
			 secretFile, err := p.client.GetSecretAttachment(strconv.Itoa(result.GetSecretResult.Secret.Id), strconv.Itoa(item.Id))
			 if err != nil {
			 	//TODO: we shouldn't skip
			 	continue
			 }
			 secretdata[strings.ToLower(item.FieldName)] = secretFile
		 } else {
			// fieldname is always lowercased.
			secretdata[strings.ToLower(item.FieldName)] = item.Value
		}
	}
	username, hasUsername := secretdata["username"]
	password, hasPassword := secretdata["password"]

	sshKey, hasSshKey := secretdata["private key"]
	sshPassphrase, _ := secretdata["private key passphrase"]

	if hasUsername && hasPassword {
		//this is a username and password secret.
		userpassSecret := new(credentials.UserPass)
		userpassSecret.Init()
		userpassSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		userpassSecret.Name = result.GetSecretResult.Secret.Name
		userpassSecret.Data = secretdata
		userpassSecret.Metadata = metadata
		userpassSecret.SetUsername(username)
		userpassSecret.SetPassword(password)
		return userpassSecret
	} else if hasSshKey{
		sshSecret := new(credentials.Ssh)
		sshSecret.Init()
		sshSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		sshSecret.Name = result.GetSecretResult.Secret.Name
		sshSecret.Data = secretdata
		sshSecret.Metadata = metadata
		sshSecret.SetKey(sshKey)
		sshSecret.SetPassphrase(sshPassphrase)
		return sshSecret
	} else {
		//this is an unknown secret type. Generic.
		genericSecret := new(credentials.Generic)
		genericSecret.Init()
		genericSecret.Id = strconv.Itoa(result.GetSecretResult.Secret.Id)
		genericSecret.Name = result.GetSecretResult.Secret.Name
		genericSecret.Data = secretdata
		genericSecret.Metadata = metadata
		return genericSecret
	}
}