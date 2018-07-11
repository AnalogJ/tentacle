package credentials

type Interface interface {
	Init()
	ToRawString() (string, error)
	ToJsonString() (string, error)
	ToTableString() (string, error)
}
