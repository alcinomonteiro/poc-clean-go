package jsonerror

import "fmt"

const JsonErrorCode string = "JSON_ERROR"

type jsonErrorType int

const (
	JsonErrorFormat jsonErrorType = iota + 1
	JsonErrorDataType
)

// JsonError -> passe Err = nill para retornar a mensagem padrão do tipo informado
type JsonError struct {
	Type jsonErrorType
	Err  error
}

func codeFormat(c int) string {
	return fmt.Sprintf("%s_%04d", JsonErrorCode, c)
}

func (jerr *JsonError) Error() string {

	code := codeFormat(int(jerr.Type))
	var msg string

	if jerr.Err != nil {
		msg = jerr.Err.Error()
	}

	switch jerr.Type {
	case JsonErrorFormat:
		if msg == "" {
			msg = "Invalid JSON Format"
		}
	case JsonErrorDataType:
		if msg == "" {
			msg = "Invalid JSON Data"
		}
	default:
		code = codeFormat(0)
		msg = "Unknown JSON error"
	}

	return fmt.Sprintf("JsonError{Code:%s,Message:%s}", code, msg)
}
