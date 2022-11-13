package betatest

/*
Status Status of the beta test.
*/
type Status string

func Status_IN_DRAFT() Status {
	return "IN_DRAFT"
}
func Status_STARTING() Status {
	return "STARTING"
}
func Status_RUNNING() Status {
	return "RUNNING"
}
func Status_STOPPING() Status {
	return "STOPPING"
}
func Status_ENDED() Status {
	return "ENDED"
}
