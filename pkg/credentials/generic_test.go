package credentials_test

import (
	"testing"

	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/stretchr/testify/require"
)

func TestGeneric_Init(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.Generic {}
	generic.Init()

	//assert
	require.Empty(t, generic.Id)
	require.Empty(t, generic.Metadata)
	require.Equal(t, generic.GetSecretType(), "generic")
}

func TestGeneric_ToJsonString(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.Generic {}
	generic.Init()
	generic.Id = "test-id"
	generic.Name = "test cred name"
	generic.Data = map[string]string{"testkey": "testvalue"}
	json, err := generic.ToJsonString()

	//assert
	require.NoError(t,err)
	require.Equal(t,`{
    "metadata": {},
    "id": "test-id",
    "name": "test cred name",
    "description": "",
    "data": {
        "testkey": "testvalue"
    }
}`, json)
}

func TestGeneric_ToRawString(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.Generic {}
	generic.Init()
	generic.Id = "test-id"
	generic.Name = "test cred name"
	generic.Data = map[string]string{"testkey": "testvalue"}
	raw, err := generic.ToRawString()

	//assert
	require.NoError(t,err)
	require.Equal(t, "id=test-id&name=test+cred+name&testkey=testvalue", raw)
}

func TestGeneric_ToYamlString(t *testing.T) {
	t.Parallel()

	//test
	generic := credentials.Generic {}
	generic.Init()
	generic.Id = "test-id"
	generic.Name = "test cred name"
	generic.Data = map[string]string{"testkey": "testvalue"}
	yml, err := generic.ToYamlString()

	//assert
	require.NoError(t,err)
	require.Equal(t,"summary:\n  metadata: {}\n  id: test-id\n  name: test cred name\n  description: \"\"\ndata:\n  testkey: testvalue\n", yml)
}