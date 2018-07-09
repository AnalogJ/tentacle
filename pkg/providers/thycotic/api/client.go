package api
//based off of https://github.com/gogosphere/gogetthychotic MIT

// https://secretserveronline.com/webservices/sswebservice.asmx API description here

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	xmld "github.com/gogosphere/xmlanswers"
)

type Client struct {
	Domain  string //"YOUADDOMAIN"
	Hosturl string //"https://yoursecretserver.yourdomain.com/SecretServer/webservices/SSWebservice.asmx"

	Username string //AD Username
	Password string //Password
}

func (c *Client)performRequest(xmlpayloadsource string, contentLengthraw int) []byte {
	contentLength := strconv.Itoa(contentLengthraw)
	client := &http.Client{}
	method := "POST"

	req, err := http.NewRequest(method, c.Hosturl, bytes.NewBuffer([]byte(xmlpayloadsource)))
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("Content-Length", contentLength)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (c *Client) List(searchTerm string) {
	uriPayLoad := xmld.WindCreds(c.Username, string(c.Password), c.Domain)
	uriPostLength := len(uriPayLoad)
	tokenxml := c.performRequest(uriPayLoad, uriPostLength)
	token := xmld.UnwindToken(tokenxml)

	searchPayload := xmld.WindSearch(token, searchTerm)
	searchPostLength := len(searchPayload)
	searchxml := c.performRequest(searchPayload, searchPostLength)
	searchName, searchID := xmld.UnwindSearch(searchxml)
	for k, v := range searchName {
		fmt.Println(searchID[k], v)
	}
	fmt.Println("END OF RESULTS: Search")

}

func (c *Client) Get(secretId string){

	uriPayLoad := xmld.WindCreds(c.Username, string(c.Password), c.Domain)
	uriPostLength := len(uriPayLoad)
	tokenxml := c.performRequest(uriPayLoad, uriPostLength)
	token := xmld.UnwindToken(tokenxml)

	tokenPayLoad := xmld.WindToken(token, secretId)
	tokenPostLength := len(tokenPayLoad)
	passxml := c.performRequest(tokenPayLoad, tokenPostLength)
	secrets := xmld.UnwindSecret(passxml)
	for _, v := range secrets {
		fmt.Println(string(v))
	}
	fmt.Println("END OF RESULTS: SecretID")
}