package api

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"io"
	"encoding/json"
)

func New(host string, token string, scheme string, httpClient *http.Client) *Client {
	c := new(Client)
	c.Host = host
	c.Token = token
	c.Scheme = scheme

	if httpClient != nil {
		c.HttpClient = httpClient
	} else {
		c.HttpClient = http.DefaultClient
	}
	return c
}

const (
	ENDPOINT_RESOURCES_LIST = "resources"
	ENDPOINT_RESOURCE_ACCOUNTS_LIST = "resources/%s/accounts"
	ENDPOINT_ACCOUNT_DETAILS = "resources/%s/accounts/%s"
	ENDPOINT_ACCOUNT_PASSWORD = "resources/%s/accounts/%s/password"
	ENDPOINT_RESOURCE_ACCOUNT_BY_PATH = "resources/getResourceIdAccountId"
)


type Client struct {
	Scheme string // "https" or "http"
	Host string //"demo.passwordmanagerpro.com:7272"
	Token string //Secret Server token

	HttpClient *http.Client
}

//func (c *Client) Test() (error) {
//
//	err := c.makeRequest("GET", ENDPOINT_RESOURCE_LIST, map[string]string{},map[string]string)
//	return err
//}

func (c *Client) GetResources() (*ResourceListResponse, error) {

	resp := new(ResourceListResponse)
	err := c.makeRequest("GET", ENDPOINT_RESOURCES_LIST, map[string]string{},map[string]string{}, resp)
	return resp, err
}

func (c *Client) GetResourceAccounts(resourceId string) (*ResourceAccountListResponse, error) {

	resp := new(ResourceAccountListResponse)
	err := c.makeRequest("GET", fmt.Sprintf(ENDPOINT_RESOURCE_ACCOUNTS_LIST, resourceId), map[string]string{},map[string]string{}, resp)
	return resp, err
}


func (c *Client) GetAccountPassword(resourceId string, accountId string) (*AccountPasswordResponse, error) {

	resp := new(AccountPasswordResponse)
	err := c.makeRequest("GET", fmt.Sprintf(ENDPOINT_ACCOUNT_PASSWORD, resourceId, accountId), map[string]string{},map[string]string{}, resp)
	return resp, err
}

func (c *Client) GetAccountDetails(resourceId string, accountId string) (*AccountDetailsResponse, error) {

	resp := new(AccountDetailsResponse)
	err := c.makeRequest("GET", fmt.Sprintf(ENDPOINT_ACCOUNT_DETAILS, resourceId, accountId), map[string]string{},map[string]string{}, resp)
	return resp, err
}

func (c *Client) GetAccountByName(resourceName string, accountName string) (*AccountByNameResponse, error) {

	resp := new(AccountByNameResponse)
	err := c.makeRequest("GET", ENDPOINT_RESOURCE_ACCOUNT_BY_PATH, map[string]string{},map[string]string{"ACCOUNTNAME": accountName, "RESOURCENAME": resourceName}, resp)
	return resp, err
}


func (c *Client) makeRequest(action string, endpoint string, data map[string]string, queryData map[string]string, resp interface{}) error {
	query := url.Values{}
	if queryData != nil {
		for key, val := range queryData {
			query.Set(key, val)
		}
	}
	query.Set("AUTHTOKEN", c.Token)

	uri := url.URL{
		Scheme:   "http",
		Host:     c.Host,
		Path:     fmt.Sprintf("restapi/json/v1/%s", endpoint),
		RawQuery: query.Encode(),
	}

	var res *http.Response
	var err error
	if action == http.MethodGet {
		res, err = c.HttpClient.Get(uri.String())
	} else {
		return fmt.Errorf("unsupport http action")
	}

	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil && err != io.EOF {
		return err
	}

fmt.Printf("\n\n\nRAW: %#v\n\n\n", string(b))

	return json.Unmarshal(b, resp)
}