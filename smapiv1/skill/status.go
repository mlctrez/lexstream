package skill

/*
Status Status of a resource.
*/
type Status string

func Status_FAILED() Status {
	return "FAILED"
}
func Status_IN_PROGRESS() Status {
	return "IN_PROGRESS"
}
func Status_SUCCEEDED() Status {
	return "SUCCEEDED"
}
