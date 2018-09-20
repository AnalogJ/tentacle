package credentials_test

import (
	"testing"

	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/stretchr/testify/require"
)

func TestBase_Init(t *testing.T) {
	t.Parallel()

	//test
	base := credentials.Summary{}
	base.Init()

	//assert
	require.Empty(t, base.Id)
	require.Empty(t, base.Metadata)
	require.Equal(t, base.GetSecretType(), "summary")
}

func TestBase_ToJsonString(t *testing.T) {
	t.Parallel()

	//test
	base := credentials.Summary{}
	base.Init()
	base.Id = "test-id"
	base.Name = "test cred name"
	json, err := base.ToJsonString()

	//assert
	require.NoError(t,err)
	require.Equal(t,json, "{\n    \"metadata\": {},\n    \"id\": \"test-id\",\n    \"name\": \"test cred name\",\n    \"description\": \"\"\n}")
}

func TestBase_ToRawString(t *testing.T) {
	t.Parallel()

	//test
	base := credentials.Summary{}
	base.Init()
	base.Id = "test-id"
	base.Name = "test cred name"
	raw, err := base.ToRawString()

	//assert
	require.NoError(t,err)
	require.Equal(t, raw, "id=test-id&name=test+cred+name")

}

