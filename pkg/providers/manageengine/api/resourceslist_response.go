package api

type ResourceListResponse struct {
	Operation ResourceList `json:"operation"`

}
type ResourceList struct {
	Name string `json:"name"`
	TotalRows int `json:"totalRows"`
	Details []ResourceItem `json:"Details"`
}

type ResourceItem struct {
	Description string `json:"RESOURCE DESCRIPTION"`
	Type string `json:"RESOURCE TYPE"`
	Id string `json:"RESOURCE ID"`
	Name string `json:"RESOURCE NAME"`
	NumberOfAccounts string `json:"NOOFACCOUNTS"`
}
