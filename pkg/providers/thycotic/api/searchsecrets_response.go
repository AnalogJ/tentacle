package api

type SearchSecretsResponse struct {
	SearchSecretsResult SearchSecretsResult
}


type SearchSecretsResult struct {
	Errors []string
	SecretSummaries []SecretSummary `xml:"SecretSummaries>SecretSummary"`
}

type SecretSummary struct {
	SecretId int
	SecretName string
	SecretTypeName string
	SecretTypeId int
	FolderId int
	IsRestricted bool
}