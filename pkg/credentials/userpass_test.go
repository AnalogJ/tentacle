package credentials_test

import (
"testing"

"github.com/analogj/tentacle/pkg/credentials"
"github.com/stretchr/testify/require"
)

func TestUserPass_Init(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.UserPass {}
	generic.Init()

	//assert
	require.Empty(t, generic.Id)
	require.Empty(t, generic.Metadata)
	require.Empty(t, generic.Data)
	require.Equal(t, generic.SecretType(), "userpass")
}

func TestUserPass_ToJsonString(t *testing.T) {
	t.Parallel()

	//test
	ssh := credentials.UserPass {}
	ssh.Init()
	ssh.Id = "test-id"
	ssh.Name = "test cred name"
	ssh.SetUsername("myusername")
	ssh.SetPassword("mypassword")
	json, err := ssh.ToJsonString()

	//assert
	require.NoError(t,err)
	require.Equal(t,`{
    "metadata": {},
    "id": "test-id",
    "name": "test cred name",
    "description": "",
    "data": {
        "password": "mypassword",
        "username": "myusername"
    }
}`, json)
}

func TestUserPass_ToRawString(t *testing.T) {
	t.Parallel()

	//test
	ssh := credentials.UserPass {}
	ssh.Init()
	ssh.Id = "test-id"
	ssh.Name = "test cred name"
	ssh.SetUsername("myusername")
	ssh.SetPassword("mypassword&")
	raw, err := ssh.ToRawString()

	//assert
	require.NoError(t,err)
	require.Equal(t, "password=mypassword%26&username=myusername", raw)
}

