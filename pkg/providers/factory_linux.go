package providers

import (
	"github.com/analogj/tentacle/pkg/utils"
	"github.com/analogj/tentacle/pkg/providers/cyberark"
	"github.com/analogj/tentacle/pkg/providers/thycotic"
	"log"
)

func Create(alias string, config interface{}) (Interface, error) {
	//stringify config keys for all provider config entries.
	config = utils.StringifyYAMLMapKeys(config)

	//begin switch for providers.
	switch providerType := config.(map[string]interface{})["type"].(string); providerType {

	//os specific providers
	//case "keychain":
	//	provider := new(keychain.provider)
	//	provider.Init(alias, config.(map[string]interface{}))
	//	return provider, nil
	//

	//alphabetical list of common providers
	case "conjur":
		return conjur.New(alias, config.(map[string]interface{}))
	case "cyberark":
		return cyberark.New(alias, config.(map[string]interface{}))
	case "lastpass":
		return lastpass.New(alias, config.(map[string]interface{}))
	case "thycotic":
		return thycotic.New(alias, config.(map[string]interface{}))

	//fall back error message
	default:
		log.Fatalf("%v type is not supported by tentacle", providerType)
		return nil, nil
	}
}