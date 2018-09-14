package api
//based off of https://github.com/gogosphere/gogetthychotic MIT

// https://secretserveronline.com/webservices/sswebservice.asmx API description here

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/tiaguinho/gosoap"
)

type Client struct {
	Domain string //"YOUADDOMAIN"
	Server string //"https://yoursecretserver.yourdomain.com/SecretServer/webservices/SSWebservice.asmx"

	//Username string //AD Username
	//Password string //Password
	Token string //Secret Server token


	//FolderCache
	FolderCache *folderNode
}

func (c *Client) Init(domain string, server string, token string) {
	c.Domain = domain
	c.Server = server
	c.Token = token

	folderCache := new(folderNode)
	folderCache.Init("-1", "root")
	c.FolderCache = folderCache
}

func (c *Client) Test() (WhoAmIResponse, error) {

	resp := WhoAmIResponse{}
	err := c.soapRequest("WhoAmI", map[string]string{"token": c.Token }, &resp)
	return resp, err
}

func (c *Client) GetFolderSubfolders(folderId string) (FolderGetAllChildrenResponse, error){

	resp := FolderGetAllChildrenResponse{}
	err := c.soapRequest("FolderGetAllChildren",
		map[string]string{
			"parentFolderId": folderId,
			"token": c.Token,
		},
		&resp)

	return resp, err
}

func (c *Client) GetFolderSecrets(folderId string) (SearchSecretsByFolderResponse, error){

	resp := SearchSecretsByFolderResponse{}
	err := c.soapRequest("SearchSecretsByFolder",
		map[string]string{
			"searchTerm": "",
			"folderId": folderId,
			"includeSubFolders": "false",
			"includeDeleted": "false",
			"includeRestricted": "false",
			"token": c.Token,
		},
		&resp)

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

func (c *Client) GetById(secretId string) (GetSecretResponse, error){

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

func (c *Client) GetByPath(secretPath string) (GetSecretResponse, error){

	secretId, err := c.getSecretIdForPath(secretPath)
	if err != nil {
		return GetSecretResponse{}, err
	}

	return c.GetById(secretId)
}



type folderNode struct {
	Id string
	Parent *folderNode
	Name string
	Children map[string]*folderNode
	Secrets map[string]string
}

func (f * folderNode) Init(id string, name string){
	f.Id = id
	f.Name = name
	f.Children = make(map[string]*folderNode)
	f.Secrets = make(map[string]string)
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

	//fmt.Printf("################################\n[DEBUG] Request %v\n%#v\nResponse Body: %#v\n###########################", method, params, string(soap.Body))

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




func (c *Client) getSecretIdForPath(secretPath string) (string, error){

	// convert slash-delimited path to list
	// first strip leading and trailing '/'
	secretPath = strings.Trim(secretPath, "/")

	secretPathComponents := strings.Split(secretPath, "/")

	secretName, folderPathComponents := secretPathComponents[len(secretPathComponents)-1], secretPathComponents[:len(secretPathComponents)-1]

	folderNode, err := c.getFolderNodeForPathComponents(folderPathComponents, c.FolderCache)
	if err != nil {
		return "", err
	}

	summaries, err := c.GetFolderSecrets(folderNode.Id)
	if err != nil {
		return "", err
	}
	//find the secret by name in summaries
	for _, summary := range summaries.SearchSecretsByFolderResult.SecretSummaries {
		if secretName == summary.SecretName{
			return strconv.Itoa(summary.SecretId), nil
		}
	}

	return "", errors.New("could not find secret by path")
}


func (c *Client) getFolderNodeForPathComponents(currentFolderPathComponents []string, parent *folderNode) (*folderNode, error){
	if len(currentFolderPathComponents) == 0{
		return parent, nil
	}

	folderName, remainingPathComponents := currentFolderPathComponents[0], currentFolderPathComponents[1:]

	if node, ok := parent.Children[folderName]; ok {
		return c.getFolderNodeForPathComponents(remainingPathComponents, node)
	} else {
		//this folder does not exist in the cache, we should attempt to retrieve it, or fail if it doesnt exist.

		parentFolderData, err := c.GetFolderSubfolders(parent.Id)
		if err != nil {
			return nil, err
		}

		for _, childFolder := range parentFolderData.FolderGetAllChildrenResult.Folders {

			childNode := new(folderNode)
			childNode.Init(strconv.Itoa(childFolder.Id), childFolder.Name)

			parent.Children[childFolder.Name] = childNode
		}

		//check if the requested folder exists now
		if node, ok := parent.Children[folderName]; ok {
			return c.getFolderNodeForPathComponents(remainingPathComponents, node)
		} else {
			return nil, errors.New(fmt.Sprintf("Could not find folder in path %v", folderName))

		}
	}



}