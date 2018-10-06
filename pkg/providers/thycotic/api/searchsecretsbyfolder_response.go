package api

type SearchSecretsByFolderResponse struct {
	SearchSecretsByFolderResult SearchSecretsByFolderResult
}


type SearchSecretsByFolderResult struct {
	Errors []string `xml:"Errors>string"`
	SecretSummaries []SecretSummary `xml:"SecretSummaries>SecretSummary"`
}

//type SecretSummary struct {
//	GetSecretId int
//	SecretName string
//	SecretTypeName string
//	SecretTypeId int
//	FolderId int
//	IsRestricted bool
//}