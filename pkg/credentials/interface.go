package credentials

type SummaryInterface interface {
	//Constructor
	Init()

	//Getters
	GetSecretType() string
	GetSecretId() string
	GetMetadata() (map[string]string)

	//Output
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToYamlString() (string, error)
 	ToTableString() (string, error)
}

type GenericInterface interface {
	//Constructor
	Init()

	//Getters
	GetSecretType() string
	GetSecretId() string
	GetData() (map[string]string)
	GetMetadata() (map[string]string)

	//Output
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToYamlString() (string, error)
	ToTableString() (string, error)
}