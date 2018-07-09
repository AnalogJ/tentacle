package api
//based off of https://github.com/hoop33/go-cyberark MIT

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

const getPasswordPath = "Accounts"

// PasswordService gets passwords
type PasswordService struct {
	client      *Client
	appID       string
	address     string
	database    string
	folder      string
	object      string
	policyID    string
	query       string
	queryFormat string
	reason      string
	safe        string
	timeout     int
	userName    string
}

// PasswordResult returns the result from getting passwords
type PasswordResult struct {
	StatusCode int
	ErrorCode  string
	ErrorMsg   string
	Content    string
	UserName   string
	Address    string
	Database   string
	PolicyID   string
	Properties map[string]string
}

func newPasswordService(client *Client) *PasswordService {
	return &PasswordService{
		client:  client,
		timeout: 30,
	}
}

// AppID sets the app ID
func (s *PasswordService) AppID(appID string) *PasswordService {
	s.appID = appID
	return s
}

// Address sets the address
func (s *PasswordService) Address(address string) *PasswordService {
	s.address = address
	return s
}

// Database sets the database
func (s *PasswordService) Database(database string) *PasswordService {
	s.database = database
	return s
}

// Folder sets the folder
func (s *PasswordService) Folder(folder string) *PasswordService {
	s.folder = folder
	return s
}

// Object sets the object
func (s *PasswordService) Object(object string) *PasswordService {
	s.object = object
	return s
}

// PolicyID sets the policy ID
func (s *PasswordService) PolicyID(policyID string) *PasswordService {
	s.policyID = policyID
	return s
}

// Query sets the query
func (s *PasswordService) Query(query string) *PasswordService {
	s.query = query
	return s
}

// QueryFormat sets the query format
func (s *PasswordService) QueryFormat(queryFormat string) *PasswordService {
	s.queryFormat = queryFormat
	return s
}

// Reason sets the reason
func (s *PasswordService) Reason(reason string) *PasswordService {
	s.reason = reason
	return s
}

// Safe sets the safe
func (s *PasswordService) Safe(safe string) *PasswordService {
	s.safe = safe
	return s
}

// Timeout sets the connection timeout
func (s *PasswordService) Timeout(timeout int) *PasswordService {
	s.timeout = timeout
	return s
}

// UserName sets the user name
func (s *PasswordService) UserName(userName string) *PasswordService {
	s.userName = userName
	return s
}

// Do runs the service
func (s *PasswordService) Do() (*PasswordResult, error) {
	if s.client == nil {
		return nil, errors.New("Client is required")
	}
	if s.appID == "" {
		return nil, errors.New("AppID is required")
	}

	path, params := s.buildURL()

	resp, err := s.client.PerformRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}
	ret := new(PasswordResult)
	if err := json.Unmarshal(resp.Body, ret); err != nil {
		return nil, err
	}
	ret.StatusCode = resp.StatusCode
	return ret, nil
}

func (s *PasswordService) buildURL() (string, url.Values) {
	params := url.Values{}

	setParam(&params, "appId", s.appID)
	setParam(&params, "address", s.address)
	setParam(&params, "database", s.database)
	setParam(&params, "folder", s.folder)
	setParam(&params, "object", s.object)
	setParam(&params, "policyID", s.policyID)
	setParam(&params, "query", s.query)
	setParam(&params, "queryFormat", s.queryFormat)
	setParam(&params, "reason", s.reason)
	setParam(&params, "safe", s.safe)
	setParam(&params, "timeout", strconv.Itoa(s.timeout))
	setParam(&params, "userName", s.userName)

	return getPasswordPath, params
}

func setParam(params *url.Values, key string, value string) {
	if value != "" {
		params.Set(key, value)
	}
}