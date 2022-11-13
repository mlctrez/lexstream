package validations

/*
ResponseValidationStatus String that specifies status of the validation. Possible values are "SUCCESSFUL" and "FAILED"
*/
type ResponseValidationStatus string

func ResponseValidationStatus_SUCCESSFUL() ResponseValidationStatus {
	return "SUCCESSFUL"
}
func ResponseValidationStatus_FAILED() ResponseValidationStatus {
	return "FAILED"
}
