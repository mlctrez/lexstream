package isp

/*
Status Current status of in-skill product.
*/
type Status string

func Status_INCOMPLETE() Status {
	return "INCOMPLETE"
}
func Status_COMPLETE() Status {
	return "COMPLETE"
}
func Status_CERTIFICATION() Status {
	return "CERTIFICATION"
}
func Status_PUBLISHED() Status {
	return "PUBLISHED"
}
func Status_SUPPRESSED() Status {
	return "SUPPRESSED"
}
