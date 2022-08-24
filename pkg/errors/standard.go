package errors

import "fmt"

var ValidationError string = "validation_error"
var ApplicationError string = "application_error"

type LagoinhaError struct {
	Message string
	Type    string
}

func (l *LagoinhaError) Error() string {
	return fmt.Sprintf("%s - %s", l.Type, l.Message)
}
