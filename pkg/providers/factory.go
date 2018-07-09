package providers

import (
	"tentacle/pkg/providers/os/darwin/keychain"
	"fmt"
	"tentacle/pkg/utils"
	"tentacle/pkg/providers/cyberark"
	"tentacle/pkg/providers/thycotic"
)

func Create(alias string, config interface{}) (Interface, error) {
	//stringify config keys for all provider config entries.
	config = utils.StringifyYAMLMapKeys(config)

	//begin switch for providers.
	switch providerType := config.(map[string]interface{})["type"].(string); providerType {
	case "keychain":
		provider := keychain.Provider {}
		provider.Init(alias, config.(map[string]interface{}))
		return &provider, nil
	case "cyberark":
		provider := cyberark.Provider {}
		provider.Init(alias, config.(map[string]interface{}))
		return &provider, nil
	case "thycotic":
		provider := thycotic.Provider {}
		provider.Init(alias, config.(map[string]interface{}))
		return &provider, nil
	default:
		fmt.Errorf("%v is not supported", providerType)
		return nil, nil
	}
}