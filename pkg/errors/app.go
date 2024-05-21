package errors

var PreferenceProviderNotFound error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "preference provider not found",
}

var RequestTimeout error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "request timed out",
}
