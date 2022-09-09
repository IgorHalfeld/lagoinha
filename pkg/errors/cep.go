package errors

var CepNotFoundError error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "cep not found",
}

var CepNotValidError error = &LagoinhaError{
	Type:    ValidationError,
	Message: "cep not valid",
}
