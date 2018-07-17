// +build darwin
package keychain

import "fmt"
import (
	goKeychain "github.com/keybase/go-keychain"
	"github.com/analogj/tentacle/pkg/utils"
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/analogj/tentacle/pkg/errors"
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

func (p *Provider) Authenticate() error {

	if location, ok := p.ProviderConfig["location"]; ok {
		//ensure that we can authenticate to this secret store.
		locationPath, _ := utils.ExpandPath(location.(string))
		k := goKeychain.NewWithPath(locationPath)

		if err := k.Status(); err == goKeychain.ErrorNoSuchKeychain {
			if p.DebugMode {
				fmt.Printf("DEBUG: no keychain found at %v for %v\n", p.ProviderConfig["location"], p.Alias)
			}
			return errors.ConfigInvalidError(fmt.Sprintf("Specified keychain does not exist for %v", p.Alias))
		} else {
			if p.DebugMode {
				fmt.Printf("DEBUG: keychain status: %#v\n", k.Status())
			}
		}
	} else if p.DebugMode {
		fmt.Println("DEBUG: No keychain path specified. Querying all.")
	}
	return nil
}

func (p *Provider) Get(queryData map[string]string) (credentials.Interface, error) {

	query := goKeychain.NewItem()
	query.SetSecClass(goKeychain.SecClassGenericPassword)
	query.SetMatchLimit(goKeychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.SetReturnData(true)

	if service, serviceOk := queryData["service"]; serviceOk {
		query.SetService(service)
	}

	if account, accountOk := queryData["account"]; accountOk {
		query.SetAccount(account)
	}

	if label, labelOk := queryData["label"]; labelOk {
		query.SetLabel(label)
	}

	if description, descriptionOk := queryData["description"]; descriptionOk {
		query.SetDescription(description)
	}

	if location, ok := p.ProviderConfig["location"]; ok {
		//ensure that we can authenticate to this secret store.
		locationPath, _ := utils.ExpandPath(location.(string))
		k := goKeychain.NewWithPath(locationPath)
		query.UseKeychain(k)
	}



	results, err := goKeychain.QueryItem(query)
	if err != nil {
		// handle error
		return nil, err
	} else if len(results) == 1 {
		if p.DebugMode {
			fmt.Println("DEBUG: found secret")
		}
		return PopulateCredential(results[0]), nil
	} else {
		if p.DebugMode {
			fmt.Println("DEBUG: no secrets found")
		}
		return nil, nil
	}
}

func (p *Provider) List(queryData map[string]string) ([]credentials.Interface, error) {
	query := goKeychain.NewItem()
	query.SetSecClass(goKeychain.SecClassGenericPassword)
	query.SetMatchLimit(goKeychain.MatchLimitAll)
	query.SetReturnAttributes(true)

	if service, serviceOk := queryData["service"]; serviceOk {
		query.SetService(service)
	}

	if account, accountOk := queryData["account"]; accountOk {
		query.SetAccount(account)
	}

	if label, labelOk := queryData["label"]; labelOk {
		query.SetLabel(label)
	}

	if description, descriptionOk := queryData["description"]; descriptionOk {
		query.SetDescription(description)
	}

	if location, ok := p.ProviderConfig["location"]; ok {
		//ensure that we can authenticate to this secret store.
		locationPath, _ := utils.ExpandPath(location.(string))
		k := goKeychain.NewWithPath(locationPath)
		query.UseKeychain(k)
	}
	results, err := goKeychain.QueryItem(query)

	if err != nil {
		// handle error
		return nil, err
	} else if len(results) >= 1 {
		if p.DebugMode {
			fmt.Println("DEBUG: found secrets")
		}

		secrets := []credentials.Interface{}
		for _, r := range results {
			secrets = append(secrets, PopulateCredential(r))
		}

		return secrets, nil
	} else {
		if p.DebugMode {
			fmt.Println("DEBUG: no secrets found")
		}
		return nil, nil
	}

	return nil, nil
}

func PopulateCredential(queryResult goKeychain.QueryResult) credentials.Interface {
	//TODO: handle non-password credentials.
	// As of now, we can only read password credentials from Keychain, so we only have to worry about password data here

	secret := new(credentials.Text)
	secret.Init()
	secret.Data = string(queryResult.Data)

	//parse metadata
	secret.Metadata["service"] = queryResult.Service
	secret.Metadata["account"] = queryResult.Account
	secret.Metadata["accessGroup"] = queryResult.AccessGroup
	secret.Metadata["label"] = queryResult.Label
	secret.Metadata["description"] = queryResult.Description

	return secret
}