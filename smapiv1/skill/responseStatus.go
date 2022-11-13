package skill

/*
ResponseStatus Status for a Response resource.
*/
type ResponseStatus string

func ResponseStatus_FAILED() ResponseStatus {
	return "FAILED"
}
func ResponseStatus_IN_PROGRESS() ResponseStatus {
	return "IN_PROGRESS"
}
func ResponseStatus_SUCCEEDED() ResponseStatus {
	return "SUCCEEDED"
}
func ResponseStatus_ROLLBACK_SUCCEEDED() ResponseStatus {
	return "ROLLBACK_SUCCEEDED"
}
func ResponseStatus_ROLLBACK_FAILED() ResponseStatus {
	return "ROLLBACK_FAILED"
}
func ResponseStatus_SKIPPED() ResponseStatus {
	return "SKIPPED"
}
