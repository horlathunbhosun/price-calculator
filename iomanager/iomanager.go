package iomanager

type IOManager interface {
	Readline() ([]string, error)
	WriteResult(data interface{}) error
}
