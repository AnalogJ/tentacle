// +build darwin
package keychain

import "fmt"
import (
	goKeychain "github.com/keybase/go-keychain"
	"log"
	"tentacle/pkg/utils"
	"tentacle/pkg/providers/base"
)

type Provider struct {
	*base.Provider
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

	if location, ok := p.providerConfig["location"]; ok {
		//ensure that we can authenticate to this secret store.
		locationPath, _ := utils.ExpandPath(location.(string))
		k := goKeychain.NewWithPath(locationPath)

		if err := k.Status(); err == goKeychain.ErrorNoSuchKeychain {
			log.Fatalf("Keychain does not exist")
		} else {
			fmt.Printf("Status: %#v\n", k.Status())
		}
	} else {
		fmt.Printf("No keychain path specified. Querying all.")
	}
	return nil
}

func (p *Provider) Get(queryData map[string]string) error {

	fmt.Printf("query data: %v", queryData)

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

	if location, ok := p.providerConfig["location"]; ok {
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
	return nil
}

func (p *Provider) List(queryData map[string]string) error {
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

	if location, ok := p.providerConfig["location"]; ok {
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
	return nil
}