package api


type GetSecretResponse struct {
	GetSecretResult GetSecretResult
}

type GetSecretResult struct {
	Errors []string `xml:"Errors>string"`
	Secret Secret
}

type Secret struct {
	Name string
	Items []SecretItem `xml:"Items>SecretItem"`
	Id int
	SecretTypeId int
	FolderId int
	Active bool
	CheckOutMinutesRemaining int
	IsCheckedOut bool
	CheckOutUserDisplayName string
	CheckOutUserId int
	IsOutOfSync bool
	IsRestricted bool
	OutOfSyncReason string
	//SecretSettings SecretSettings
	//SecretPermissions SecretPermissions
}

type SecretItem struct {
	Id int
	FieldDisplayName string
	FieldName string
	IsFile bool
	IsNotes bool
	IsPassword bool
	Value string
}
