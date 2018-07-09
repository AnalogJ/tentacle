package credentials

type UserPass struct {
	*Base
	Username string
	Password string
}