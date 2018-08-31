package providers

import (
	"github.com/analogj/tentacle/pkg/credentials"
	"gopkg.in/urfave/cli.v2"
)

// Create mock using:
// mockgen -source=pkg/engine/interface.go -destination=pkg/engine/mock/mock_engine.go
type Interface interface {

	Command() *cli.Command

	/*
	This function should initialize the class and run any basic validation checks of the environment/system.

	Ie. ensure that required software is installed (eg. Java, CLI's) and that the file system is configured as
	expected.
	 */
	Init(alias string, config map[string]interface{}) error

	/*
	This function should attempt to authenticate against the credential management service with the provided auth
	credentials.

	Ie. using username/password/api token
	 */
	Authenticate() error

	/*
	This function should attempt to retrieve a credential specified via queryData
	 */
	Get(queryData map[string]string) (credentials.Interface, error)

	/*
	This function should attempt to retrieve a list of credentials available.
	This may not be available in all providers.
	 */
	List(queryData map[string]string) ([]credentials.Interface, error)
}