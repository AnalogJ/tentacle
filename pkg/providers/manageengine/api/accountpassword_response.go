package api

type AccountPasswordResponse struct {
	Operation AccountPassword `json:"operation"`

}
type AccountPassword struct {
	Name string `json:"name"`
	Details struct {
		Password string `json:"PASSWORD"`
	} `json:"Details"`
}