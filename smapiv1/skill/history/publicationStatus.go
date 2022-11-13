package history

/*
PublicationStatus The publication status of the skill when this interaction occurred
*/
type PublicationStatus string

func PublicationStatus_Development() PublicationStatus {
	return "Development"
}
func PublicationStatus_Certification() PublicationStatus {
	return "Certification"
}
