package skill

/*
Reason The reason to withdraw.
*/
type Reason string

func Reason_TEST_SKILL() Reason {
	return "TEST_SKILL"
}
func Reason_MORE_FEATURES() Reason {
	return "MORE_FEATURES"
}
func Reason_DISCOVERED_ISSUE() Reason {
	return "DISCOVERED_ISSUE"
}
func Reason_NOT_RECEIVED_CERTIFICATION_FEEDBACK() Reason {
	return "NOT_RECEIVED_CERTIFICATION_FEEDBACK"
}
func Reason_NOT_INTEND_TO_PUBLISH() Reason {
	return "NOT_INTEND_TO_PUBLISH"
}
func Reason_OTHER() Reason {
	return "OTHER"
}
