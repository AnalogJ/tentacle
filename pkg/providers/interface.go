package providers

import (
	"github.com/analogj/tentacle/pkg/credentials"
	"gopkg.in/urfave/cli.v2"
	"net/http"
)


/*
	This function should initialize the class and run any basic validation checks of the environment/system.

	Ie. ensure that required software is installed (eg. Java, CLI's) and that the file system is configured as
	expected.

	New should not make any network calls
 */
//New(alias string, config map[string]interface{}) (*provider, error)


// Create mock using:
// mockgen -source=pkg/engine/interface.go -destination=pkg/engine/mock/mock_engine.go
type Interface interface {

	Command() *cli.Command




	/*
	This function should return a map of capabilites (set to true) that this provider supports.
	Available capabilites are as follows:

	`list` - specify if the `.List` function is defined for this provider.
	`get` - specify if the `.Get` function is defined for this provider
	`get_by_id` - specify the `.Get` function can be used with just an `id` property
	`get_by_path` - specify the `.Get` function can be used with just an `path` property
	`cred_userpass` - specify that the `.Get` function can return `UserPass` type credentials
	`cred_text` - specify that the `.Get` function can return `Text` type credentials
	`cred_ssh` - specify that the `.Get` function can return `SSH` type credentials
	*/
	Capabilities() map[string]bool

	/*
	This function should attempt to authenticate against the credential management service with the provided auth
	credentials.

	Ie. using username/password/api token
	 */
	Authenticate() error

	/*
	This function should attempt to retrieve a credential specified via queryData
	 */
	Get(queryData map[string]string) (credentials.GenericInterface, error)

	/*
	This function should attempt to retrieve a list of credentials available.
	Can not contain any secret data (.Data) only metadata (.Metadata)
	This may not be available in all providers.
	 */
	List(queryData map[string]string) ([]credentials.SummaryInterface, error)


	/*
	This function is defined in the base provider and is used to override the http.Client class used by the providers.
	This is required for HTTP proxy support & go-vcr replay during testing
	*/
	SetHttpClient(httpClient *http.Client)
}