package api

type ResourceAccountListResponse struct {
	Operation ResourceAccountList `json:"operation"`

}
type ResourceAccountList struct {
	Name string `json:"name"`
	TotalRows int `json:"totalRows"`
	Details struct {
		ResourceId string `json:"RESOURCE ID"`
		ResourceName string `json:"RESOURCE NAME"`
		ResourceType string `json:"RESOURCE TYPE"`
		ResourceDescription string `json:"RESOURCE DESCRIPTION"`
		Accounts []AccountItem `json:"ACCOUNT LIST"`
	} `json:"Details"`
}

type AccountItem struct {
	Id string `json:"ACCOUNT ID"`
	Name string `json:"ACCOUNT NAME"`
	PasswordId string `json:"PASSWDID"`
}
