package api
//based off of https://github.com/gogosphere/gogetthychotic MIT

// https://secretserveronline.com/webservices/sswebservice.asmx API description here

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/tiaguinho/gosoap"
)

type Client struct {
	Domain string //"YOUADDOMAIN"
	Server string //"https://yoursecretserver.yourdomain.com/SecretServer/webservices/SSWebservice.asmx"

	//Username string //AD Username
	//Password string //Password
	Token string //Secret Server token
}


func (c * Client) soapRequest(method string, params map[string]string, response interface{}) (error){
	soap, err := gosoap.SoapClient(fmt.Sprintf("https://%s/webservices/sswebservice.asmx?WSDL", c.Server))
	if err != nil {
		return err
	}

	err = soap.Call(method, params)
	if err != nil {
		return err
	}

	err = soap.Unmarshal(response)
	if err != nil {
		return err
	}
	return nil
}


func (c *Client)performRequest(xmlpayloadsource string, contentLengthraw int) ([]byte, error) {
	contentLength := strconv.Itoa(contentLengthraw)
	client := &http.Client{}
	method := "POST"

	req, err := http.NewRequest(method, fmt.Sprintf("https://%s/webservices/sswebservice.asmx", c.Server), bytes.NewBuffer([]byte(xmlpayloadsource)))
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("Content-Length", contentLength)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Test() (WhoAmIResponse, error) {

	resp := WhoAmIResponse{}
	err := c.soapRequest("WhoAmI", map[string]string{"token": c.Token }, &resp)
	return resp, err
}


func (c *Client) List(searchTerm string) (SearchSecretsResponse, error) {

	resp := SearchSecretsResponse{}
	err := c.soapRequest("SearchSecrets",
		map[string]string{
			"includeDeleted": "0",
			"includeRestricted": "1",
			"searchTerm": searchTerm,
			"token": c.Token,
		},
		&resp)
	return resp, err

}

func (c *Client) Get(secretId string) (GetSecretResponse, error){

	resp := GetSecretResponse{}
	err := c.soapRequest("GetSecret",
		map[string]string{
			"loadSettingsAndPermissions": "0",
			"secretId": secretId,
			"token": c.Token,
		},
		&resp)

	return resp, err
}