package credentials

type Base struct {
	Metadata map[string]string
	Type string
}

func (b *Base) Init() {
	//do nothing
	b.Metadata = map[string]string{}
	b.Type = ""
}
