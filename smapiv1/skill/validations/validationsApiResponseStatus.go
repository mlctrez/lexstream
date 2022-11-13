package validations

/*
ValidationsApiResponseStatus String that specifies the current status of validation execution. Possible values are "IN_PROGRESS", "SUCCESSFUL", and "FAILED".
*/
type ValidationsApiResponseStatus string

func ValidationsApiResponseStatus_IN_PROGRESS() ValidationsApiResponseStatus {
	return "IN_PROGRESS"
}
func ValidationsApiResponseStatus_SUCCESSFUL() ValidationsApiResponseStatus {
	return "SUCCESSFUL"
}
func ValidationsApiResponseStatus_FAILED() ValidationsApiResponseStatus {
	return "FAILED"
}
