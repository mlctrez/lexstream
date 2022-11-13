package skill

/*
VersionSubmissionStatus The lifecycle status of the skill version submission.
* `LIVE` - The skill version is in the live stage
* `CERTIFIED` - The skill version has gone through the certification review process and has been certified.
* `IN_REVIEW` - The skill version is currently under review for certification and publication. During this time, you cannot edit the configuration.
* `FAILED_CERTIFICATION` - The skill version has been submitted for certification, however it has failed certification review. Please submit a new version for certification.
* `HIDDEN` - The skill version has been published but is no longer available to new users for activation. Existing users can still invoke this skill if it is the most recent version.
* `REMOVED` - The skill version has been published but removed for use, due to Amazon's policy violation. You can update your skill and publish a new version to live to address the policy violation.
* `WITHDRAWN_FROM_CERTIFICATION` - The skill version was submitted for certification but was withdrawn from review.
*/
type VersionSubmissionStatus string

func VersionSubmissionStatus_LIVE() VersionSubmissionStatus {
	return "LIVE"
}
func VersionSubmissionStatus_CERTIFIED() VersionSubmissionStatus {
	return "CERTIFIED"
}
func VersionSubmissionStatus_IN_REVIEW() VersionSubmissionStatus {
	return "IN_REVIEW"
}
func VersionSubmissionStatus_FAILED_CERTIFICATION() VersionSubmissionStatus {
	return "FAILED_CERTIFICATION"
}
func VersionSubmissionStatus_HIDDEN() VersionSubmissionStatus {
	return "HIDDEN"
}
func VersionSubmissionStatus_REMOVED() VersionSubmissionStatus {
	return "REMOVED"
}
func VersionSubmissionStatus_WITHDRAWN_FROM_CERTIFICATION() VersionSubmissionStatus {
	return "WITHDRAWN_FROM_CERTIFICATION"
}
