package credentials

type BaseInterface interface {
	SecretType() string

	Init()
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}

type GenericInterface interface {
	SecretType() string

	Init()
	GetData() (map[string]string)
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}
