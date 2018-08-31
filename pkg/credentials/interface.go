package credentials

type Interface interface {
	SecretType() string

	Init()
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}