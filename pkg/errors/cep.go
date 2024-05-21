package errors

var CepNotFoundError error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "cep not found",
}

var TooManyRequestsError error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "too many requests error",
}

var InternalServerError error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "internal server error",
}

var CepNotValidError error = &LagoinhaError{
	Type:    ValidationError,
	Message: "cep not valid",
}
