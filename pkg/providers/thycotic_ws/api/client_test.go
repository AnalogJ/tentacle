package api_test

import (
	"testing"

	"github.com/analogj/tentacle/pkg/providers/thycotic_ws/api"
	"github.com/stretchr/testify/require"
)

func TestClient_GetFolderSubfolders(t *testing.T) {
	t.Parallel()

	//test
	client := new(api.Client)
	client.Init("YOUR_AD_DOMAIN", "", "")
	resp, err := client.GetFolderSubfolders("-1")


	t.Logf("GET FOLDER ID RESPONSE: %#v", resp)

	//assert
	require.NoError(t,err)
	//require.Equal(t, raw, "id=test-id&name=test+cred+name")

}

func TestClient_GetFolderSecrets(t *testing.T) {
	t.Parallel()

	//test
	client := new(api.Client)
	client.Init("YOUR_AD_DOMAIN", "", "")
	resp, err := client.GetFolderSecrets("-1")


	t.Logf("GET FOLDER SECRETS RESPONSE: %#v", resp)

	//assert
	require.NoError(t,err)
	//require.Equal(t, raw, "id=test-id&name=test+cred+name")

}