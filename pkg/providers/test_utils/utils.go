package test_utils

import (
	"testing"
	"net/http"
	"crypto/tls"
	"path"
	"strings"
	"github.com/seborama/govcr"
	"net/http/cookiejar"
)

//TODO: set to true when development is complete and no new recordings need to be created (CI mode enabled)
const DISABLE_RECORDINGS = false


func ProviderVcrSetup(t *testing.T) *http.Client {
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, //disable certificate validation because we're playing back http requests.
	}

	jar, _ := cookiejar.New(nil)
	insecureClient := http.Client{
		Transport: tr,
		Jar: jar,
	}

	vcrConfig := govcr.VCRConfig{
		Logging:      true,
		CassettePath: path.Join("testdata", "govcr-fixtures"),
		Client:       &insecureClient,
		ExcludeHeaderFunc: func(key string) bool {
			// HTTP headers are case-insensitive
			return strings.ToLower(key) == "user-agent" || strings.ToLower(key) == "authorization"
		},
		RequestFilterFunc: func(reqHeader http.Header, reqBody []byte) (*http.Header, *[]byte) {
			reqHeader.Set("Authorization", "Basic UExBQ0VIT0xERVI6UExBQ0VIT0xERVI=") //placeholder:placeholder

			return &reqHeader, &reqBody
		},

		//this line ensures that we do not attempt to create new recordings.
		//Comment this out if you would like to make recordings.
		DisableRecording: DISABLE_RECORDINGS,
	}

	vcr := govcr.NewVCR(t.Name(), &vcrConfig)
	return vcr.Client
}