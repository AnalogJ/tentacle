package thycotic_test

import (
	"testing"
	"github.com/analogj/tentacle/pkg/validator"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/stretchr/testify/suite"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestThycoticProvider_TestSuite(t *testing.T) {
	testSuite := new(validator.ProviderTestSuite)

	testSuite.ProviderConfig = map[string]interface{}{
		"type": "thycotic",
		"domain": "",
		"server": "tentacletest.secretservercloud.com",
		"token": "AgKc2oenOWcuNq_3zEdljpko53QMl5d64jMtzielAVri9kY5JkNsukfBQbEmSEKVRll5ThDcoMtvJEs8ACJWLvtDCPZi2YoFQRVIkONBrafRkJFImbTdMIcrChy8WL_3kQup-8rA12qegkNwUR8HTRSUtsru6YdkXbPT6fSyW5VMbCo_U_Mr59Ur7uyfwc4EPe3uWU2kVlUbAyW3GXaz202C-YA51Hc4wJ1Rx9tUN5exSlR9aKusgYG96LEhrfFJ9CwyEhZnYzzPRtxoADDuIUhJLrf_cGjApJ945HwEUqecmplCR_5aSIrpjpgllAExUhxR0LGHaGO7tduuGdlXopkRvyuaFjr3prDk_FHiPFR-Ffj6BjSoNlzszfnL-UqtbrjG1p4pxK5D5PG50nVHGxgeAc9yZioI6peNBjjNZ_xh_awaKieq-eBe-9S9aJoKy9WsXpoHnH5reChIizSCTQ2FtAAHkg-Ea6Dijkx2r2LHjNPqZPBAJF156V9nNTmQgPMqVwH4vehXSSPKufQO7Vy8DHrVIxE7jUBelePX2Y8gg2iMOjjPqMgjDNow8alO6uQ-d6iMRjvkRwjLgJmxgs70RD-thT2KCKcOHA8SxoGF80JknwHpO2hNIG4qkPEi1qo", //junk password only used for testing. please don't change it :(
	}
	testSuite.Get_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "1"},
		Data:      map[string]string{"password":"suzXFPeiAXU8CU7@8st&", "server":"", "username":"testusername"},
		Metadata:  map[string]string{"active":"true", "folderid":"2", "secrettypeid":"2", "notes":"userpass notes"},
	}

	testSuite.GetById_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "2"},
		Data:      map[string]string{"server":"resource info", "username":"test", "password":"8pX*zUdjOW^xcveoj@Cd"},
		Metadata:  map[string]string{"active":"true", "folderid":"3", "secrettypeid":"2", "notes":"example.com test note"},
	}

	testSuite.GetByPath_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"path": "/Personal Folders/tentacle@mailinator.com/subfolder/example.com"},
		Data:      map[string]string{"server":"resource info", "username":"test", "password":"8pX*zUdjOW^xcveoj@Cd"},
		Metadata:  map[string]string{"active":"true", "folderid":"3", "secrettypeid":"2", "notes":"example.com test note"},
	}

	testSuite.Get_UserPass_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "2"},
		Data:      map[string]string{"server":"resource info", "username":"test", "password":"8pX*zUdjOW^xcveoj@Cd"},
		Metadata:  map[string]string{"secrettypeid":"2", "notes":"example.com test note", "active":"true", "folderid":"3"},
	}

	testSuite.Get_Ssh_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "3"},
		Data:      map[string]string{"private key passphrase":"", "public key":"", "private key":"-----BEGIN RSA PRIVATE KEY-----\r\nMIIEpQIBAAKCAQEA10k2lpecat1CtsbzbAUHNiGbkEFo2xgw0VYHYOioxkSRitR1\r\ndq6dH/J0SsDTLm3DQOgdsrpU/OzEeX4m0IZyre/Gpik3ehocL+jzLpRn5SxwcONv\r\nY2Pv/7a+kC9nPveRv+EpsEWwJFMjRr9GAGMp8Agshr/5+qbKl298IlB0Dh4azJ6w\r\nIarQGr2i9NbuA0T+R78RPJEenvMcA+CsW0+oT3D5BkC62ikTH/qM+nnYSqHfjQNl\r\nbx0qj2zcw02ERp9xbgqkq7EOrmnwQf6oAIARtVk/xwd1NelQ9IFV1GceaNIlSYiC\r\n+Vc25EW9H3/0p68gm6jj9VmB3EBH7GlzLPKxIQIEAAEAAQKCAQEA1n20LH+rMPF3\r\neYFon3O3BqCfTC9PGBLW+In82fmTxc4lL6uTyIYaN+0pHh1VikmDopRgmmR5LnF7\r\nIwykAVgiw9dEFOK1yipfcQBn4T2s8bC/6kk+/khgomZhIGiKNjsUdJcSIPSdlONc\r\nHy2MwfAKpYKPUkXM+oTZDd3eUJXVuwVBeFdy/W0KkbrF6S32FUrqCJIak1bI6//1\r\nJcVE/4FGqZ/aC+StWM/jFGAHerOB8qoJ6gJ83zK+69o/P4ks/zMyO4hlqohif+rp\r\nmLKw5rsWvELOr5/PllHHUVCEGe7YgExCJUpJd42UznoaRlSU/+J5vmlKdKhNPJ2K\r\nTKLbk6xuvQKBgQDkgkmHT838UjvpsY3l0rb+bn8RqhctMgnNc80BJud2qZbAkYbD\r\nHs+hZQ+U731wonSTsZ9ilhAd9YaoOCJxgZhw1tsNAJ0hMX7asM5/OV7Ust8AV4IM\r\n9Pr7x26mru8lHmlF2hcdoL1xOoY/AkR7e9rQuPrTaSX3GCG9rEWI3YMKBwKBgQDx\r\nL64mQ5TaB53LDBMvPA7Uzj6AA/NogzxxrmjOlfC9BilfTIxrnPOX9KQzkUra3HWH\r\nR0gtw8RlFZl9OM4cuQWBwAxuXkSRPscsPu1uXDaYpmYFcn+2rqI96BLk7aZ6owiL\r\nTrIiiN0ovfadQR1lht6K41amY1BK8B0uDiu3cQRBlwKBgCEjQ8Z7rEOLKWUaepl5\r\nlVAQdhz5raLAPuusf25LVgW/Uj1a3VAuh63AGiJfGLHc3UsN1y9U08GEeaKrgVM0\r\nmAbFIb8g96h4pQzR1yBcEYSG7BAAoLuAS2V35nQFqmiXoGyg0/lX9iEVe6JnkcbW\r\nj0T1jmpaNUdAGKSI9wyCEx7BAoGBAIjIscuhqFm+2A6/2mF/finSjj+r/e/X/f+7\r\nGlWPU+jgBba2gyzE5qHXbQiR7hIR1dXS52yNCmUJyvLptHs1s5nSgTpW3CxDnlHJ\r\nnn5obc5FELcmKKhCgmD5rT2ISlJjBAV9rClJ1aO1mJe2xiU+SBgctpfG21KkuufT\r\nwZdX0UIhAoGAcjRpVyE9nXh2eFDTKRHi/LX+Oz0wF6hY7O9YsznNZrMRZRiargNq\r\nYOZMMfboDWAnnpBSjQ3EwIVQRJPl0VLIPL4EL/C/g7LlPsHcrNfXTD15OE7IIIq9\r\nId50nLBDbPeQK++XAyV4AmktMc3BaRM18Bmq8qOPnsPw8YBM6tmMX1Q=\r\n-----END RSA PRIVATE KEY-----\r\n"},
		Metadata:  map[string]string{"notes":"test ssh key pair", "active":"true", "folderid":"2", "secrettypeid":"6026"},
	}

	testSuite.Get_Text_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "5"},
		Data:      map[string]string{},
		Metadata:  map[string]string{},
	}

	var summaryList []credentials.SummaryInterface
	summary1 := new(credentials.Summary)
	summary1.Init()
	summary1.Id = "5"
	summary1.Name = "combination lock"
	summary1.Metadata = map[string]string {
		"folderid": "2",
		"secrettypeid": "4",
	}

	summary2 := new(credentials.Summary)
	summary2.Init()
	summary2.Id = "2"
	summary2.Name = "example.com"
	summary2.Metadata = map[string]string {
		"folderid": "3",
		"secrettypeid": "2",
	}

	summary3 := new(credentials.Summary)
	summary3.Init()
	summary3.Id = "3"
	summary3.Name = "ssh key"
	summary3.Metadata = map[string]string {
		"folderid": "2",
		"secrettypeid": "6026",
	}

	summary4 := new(credentials.Summary)
	summary4.Init()
	summary4.Id = "1"
	summary4.Name = "test userpass"
	summary4.Metadata = map[string]string {
		"folderid": "2",
		"secrettypeid": "2",
	}

	summaryList = append(summaryList, summary1)
	summaryList = append(summaryList, summary2)
	summaryList = append(summaryList, summary3)
	summaryList = append(summaryList, summary4)

	testSuite.List_TestData = validator.ListRequestTestData{
		QueryData: map[string]string{},
		Results: summaryList,
	}

	suite.Run(t, testSuite)
}