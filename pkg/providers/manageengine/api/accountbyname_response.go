package api

type AccountByNameResponse struct {
	Operation AccountByName `json:"operation"`

}
type AccountByName struct {
	Name string `json:"name"`
	Details struct {
		ResourceId string `json:"RESOURCEID"`
		AccountId string `json:"ACCOUNTID"`
	} `json:"Details"`
}