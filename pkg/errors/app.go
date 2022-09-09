package errors

var PreferenceProviderNotFound error = &LagoinhaError{
	Type:    ApplicationError,
	Message: "preference provider not found",
}
