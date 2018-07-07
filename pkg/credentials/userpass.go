package credentials

type UserPass struct {
	*Secret
	Username string
	Password string
}