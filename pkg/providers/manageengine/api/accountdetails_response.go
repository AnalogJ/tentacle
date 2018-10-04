package api

type AccountDetailsResponse struct {
	Operation AccountDetails `json:"operation"`

}
type AccountDetails struct {
	Name string `json:"name"`
	Details struct {
		Description string `json:"DESCRIPTION"`
	} `json:"Details"`
}