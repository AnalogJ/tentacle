package credentials_test

import (
"testing"

"github.com/analogj/tentacle/pkg/credentials"
"github.com/stretchr/testify/require"
)

func TestSsh_Init(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.Ssh {}
	generic.Init()

	//assert
	require.Empty(t, generic.Id)
	require.Empty(t, generic.Metadata)
	require.Equal(t, generic.SecretType(), "ssh")
}

func TestSsh_ToJsonString(t *testing.T) {
	t.Parallel()

	//test
	ssh := credentials.Ssh {}
	ssh.Init()
	ssh.Id = "test-id"
	ssh.Name = "test cred name"
	ssh.SetKey("multiline\nkey\nvalue")
	json, err := ssh.ToJsonString()

	//assert
	require.NoError(t,err)
	require.Equal(t,`{
    "metadata": {},
    "id": "test-id",
    "name": "test cred name",
    "description": "",
    "data": {
        "key": "multiline\nkey\nvalue"
    }
}`, json)
}

func TestSsh_ToRawString(t *testing.T) {
	t.Parallel()

	//test
	ssh := credentials.Ssh {}
	ssh.Init()
	ssh.Id = "test-id"
	ssh.Name = "test cred name"
	ssh.SetKey("multiline\nkey\nvalue")
	raw, err := ssh.ToRawString()

	//assert
	require.NoError(t,err)
	require.Equal(t, "multiline\nkey\nvalue", raw)
}

