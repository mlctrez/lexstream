package skill

/*
PublicationStatus Publication status of the skill.
It is associated with the skill's stage. Skill in 'development' stage can have publication status as 'DEVELOPMENT' or 'CERTIFICATION'. Skill in 'certified' stage can have publication status as 'CERTIFIED'. 'Skill in 'live' stage can have publication status as 'PUBLISHED', 'HIDDEN' or 'REMOVED'.
* `DEVELOPMENT` - The skill is available only to you. If you have enabled it for testing, you can test it on devices registered to your developer account.
* `CERTIFICATION` - Amazon is currently reviewing the skill for publication. During this time, you cannot edit the configuration.
* `CERTIFIED` - The skill has been certified and ready to be published. Skill can be either published immediately or an future release date can be set for the skill. You cannot edit the configuration for the certified skills. To start development, make your changes on the development version.
* `PUBLISHED` - The skill has been published and is available to users. You cannot edit the configuration for live skills. To start development on an updated version, make your changes on the development version instead.
* `HIDDEN` - The skill has been published but is no longer available to new users for activation. Existing users can still invoke this skill.
* `REMOVED` - The skill has been published but removed for use, due to Amazon's policy violation. You can update your skill and publish a new version to live to address the policy violation.
*/
type PublicationStatus string

func PublicationStatus_DEVELOPMENT() PublicationStatus {
	return "DEVELOPMENT"
}
func PublicationStatus_CERTIFICATION() PublicationStatus {
	return "CERTIFICATION"
}
func PublicationStatus_CERTIFIED() PublicationStatus {
	return "CERTIFIED"
}
func PublicationStatus_PUBLISHED() PublicationStatus {
	return "PUBLISHED"
}
func PublicationStatus_HIDDEN() PublicationStatus {
	return "HIDDEN"
}
func PublicationStatus_REMOVED() PublicationStatus {
	return "REMOVED"
}
