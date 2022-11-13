package evaluations

type Status string

func Status_PASSED() Status {
	return "PASSED"
}
func Status_FAILED() Status {
	return "FAILED"
}
func Status_IN_PROGRESS() Status {
	return "IN_PROGRESS"
}
func Status_ERROR() Status {
	return "ERROR"
}
