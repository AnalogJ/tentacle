package providers

import (
	"gopkg.in/urfave/cli.v2"
	"github.com/analogj/tentacle/pkg/credentials"
)

// Create mock using:
// mockgen -source=pkg/engine/interface.go -destination=pkg/engine/mock/mock_engine.go
type Interface interface {
	Init(alias string, config map[string]interface{}) error
	Command() *cli.Command
	Authenticate() error
	Get(queryData map[string]string) (credentials.Interface, error)
	List(queryData map[string]string) ([]credentials.Interface, error)
}