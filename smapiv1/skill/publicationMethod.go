package skill

/*
PublicationMethod Determines if the skill should be submitted only for certification and manually publish later or publish immediately after the skill is certified. Omitting the publication method will default to auto publishing.
*/
type PublicationMethod string

func PublicationMethod_MANUAL_PUBLISHING() PublicationMethod {
	return "MANUAL_PUBLISHING"
}
func PublicationMethod_AUTO_PUBLISHING() PublicationMethod {
	return "AUTO_PUBLISHING"
}
