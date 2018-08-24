package api

type WhoAmIResponse struct {
	WhoAmIResult UserInfoResult
}

type UserInfoResult struct {
	UserId int
	DisplayName string
	DomainId int
	DomainName string
	Errors []string
	KnownAs string
}
