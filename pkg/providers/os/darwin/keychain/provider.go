// +build darwin
package keychain

import "fmt"
import (
	goKeychain "github.com/keybase/go-keychain"
	"log"
)

type Provider struct {
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
	query.SetService(p.providerConfig["service"].(string))
	query.SetAccount(queryData["id"])
	query.SetAccessGroup(p.providerConfig["access_group"].(string))
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
	query.SetService(p.providerConfig["service"].(string))
	query.SetAccount(queryData["id"])
	query.SetAccessGroup(p.providerConfig["access_group"].(string))
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