package credentials

type SummaryInterface interface {
	//Constructor
	Init()

	//Getters
	GetSecretType() string
	GetSecretId() string
	GetMetaData() (map[string]string)

	//Output
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}

type GenericInterface interface {
	//Constructor
	Init()

	//Getters
	GetSecretType() string
	GetSecretId() string
	GetData() (map[string]string)
	GetMetaData() (map[string]string)

	//Output
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}