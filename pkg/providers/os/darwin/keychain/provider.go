// +build darwin
package keychain

import "fmt"
import (
	goKeychain "github.com/keybase/go-keychain"
	"log"
)

type Provider struct {
	providerConfig ProviderConfig
}

func (p *Provider) Authenticate() error {
	//ensure that we can authenticate to this secret store.
	k := goKeychain.NewWithPath("System.keychain")
	if err := k.Status(); err == goKeychain.ErrorNoSuchKeychain {
		log.Fatalf("Keychain doesn't exist yet")
	} else {
		fmt.Printf("Status: %#v\n", k.Status())

	}
	return nil
}

func (p *Provider) Get(queryData map[string]string) error {
	query := goKeychain.NewItem()
	query.SetSecClass(goKeychain.SecClassGenericPassword)
	query.SetService(p.providerConfig.Service)
	query.SetAccount(queryData["id"])
	query.SetAccessGroup(p.providerConfig.AccessGroup)
	query.SetMatchLimit(goKeychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.UseKeychain(goKeychain.NewWithPath(""))
	results, err := goKeychain.QueryItem(query)
	if err != nil {
		// handle error
		fmt.Printf("%#v\n", err)
	} else {
		fmt.Print("Found Results!\n")
		for _, r := range results {
			fmt.Printf("%#v\n", r)
		}
	}
	return nil
}

func (p *Provider) List(queryData map[string]string) error {
	query := goKeychain.NewItem()
	query.SetSecClass(goKeychain.SecClassGenericPassword)
	query.SetService(p.providerConfig.Service)
	query.SetAccount(queryData["id"])
	query.SetAccessGroup(p.providerConfig.AccessGroup)
	query.SetMatchLimit(goKeychain.MatchLimitAll)
	query.SetReturnAttributes(true)
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