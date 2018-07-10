// +build darwin
package keychain

import "fmt"
import (
	goKeychain "github.com/keybase/go-keychain"
	"tentacle/pkg/utils"
	"tentacle/pkg/providers/base"
	"tentacle/pkg/credentials"
	"tentacle/pkg/errors"
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
		fmt.Printf("%#v\n", err)
	} else {
		fmt.Print("Found Results!\n")
		for _, r := range results {
			fmt.Printf("%#v\n", r)

			fmt.Printf( "secret: %v" ,string(r.Data))

		}
	}
	return nil, nil
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
	} else {
		for _, r := range results {
			fmt.Printf("%#v\n", r)
		}
	}
	return nil, nil
}