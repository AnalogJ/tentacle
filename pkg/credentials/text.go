package credentials

type Text struct {
	*Secret
	Data string
}