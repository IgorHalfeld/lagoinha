package errors

func HttpError(statusCode int) error {
	switch statusCode {
	case 429:
		return TooManyRequestsError
	case 500:
		return InternalServerError
	default:
		return CepNotFoundError
	}
}
