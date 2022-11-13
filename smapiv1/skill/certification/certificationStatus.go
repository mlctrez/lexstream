package certification

/*
CertificationStatus String that specifies the current status of skill's certification Possible values are "IN_PROGRESS", "SUCCEEDED", "FAILED" and "CANCELLED"
*/
type CertificationStatus string

func CertificationStatus_IN_PROGRESS() CertificationStatus {
	return "IN_PROGRESS"
}
func CertificationStatus_SUCCEEDED() CertificationStatus {
	return "SUCCEEDED"
}
func CertificationStatus_FAILED() CertificationStatus {
	return "FAILED"
}
func CertificationStatus_CANCELLED() CertificationStatus {
	return "CANCELLED"
}
