package publication

/*
SkillPublicationStatus Status of publishing
*/
type SkillPublicationStatus string

func SkillPublicationStatus_IN_PROGRESS() SkillPublicationStatus {
	return "IN_PROGRESS"
}
func SkillPublicationStatus_SUCCEEDED() SkillPublicationStatus {
	return "SUCCEEDED"
}
func SkillPublicationStatus_FAILED() SkillPublicationStatus {
	return "FAILED"
}
func SkillPublicationStatus_CANCELLED() SkillPublicationStatus {
	return "CANCELLED"
}
func SkillPublicationStatus_SCHEDULED() SkillPublicationStatus {
	return "SCHEDULED"
}
