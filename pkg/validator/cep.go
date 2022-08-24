package validator

func ExceedsCepMaximumSize(cepRaw string) (status bool) {
	const cepSize = 8
	cepLength := len(cepRaw)
	if cepLength <= cepSize {
		status = false
	} else {
		status = true
	}
	return status
}
