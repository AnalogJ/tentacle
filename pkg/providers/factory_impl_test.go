package providers

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/analogj/tentacle/pkg/providers/conjur"
	"github.com/analogj/tentacle/pkg/providers/cyberark"
	"github.com/analogj/tentacle/pkg/providers/lastpass"
	"github.com/analogj/tentacle/pkg/providers/thycotic"
)

func TestConjurProvider(t *testing.T) {
	eng := new(conjur.Provider)
	require.Implements(t, (*Interface)(nil), eng, "should implement the Provider interface")
}

func TestCyberarkProvider(t *testing.T) {
	eng := new(cyberark.Provider)
	require.Implements(t, (*Interface)(nil), eng, "should implement the Provider interface")
}

func TestLastpassProvider(t *testing.T) {
	eng := new(lastpass.Provider)
	require.Implements(t, (*Interface)(nil), eng, "should implement the Provider interface")
}

func TestThycoticProvider(t *testing.T) {
	eng := new(thycotic.Provider)
	require.Implements(t, (*Interface)(nil), eng, "should implement the Provider interface")
}
