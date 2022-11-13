package skill

/*
CloneLocaleRequestStatus Status for a locale clone request
CloneLocale API initiates cloning from a source locale to all target locales.
  - `IN_PROGRESS` status would indicate the clone is still in progress.
  - `SUCCEEDED` status would indicate the source locale was cloned successfully to all target locales.
  - `INELIGIBLE` status would indicate the source locale was ineligible to be cloned to all target locales.
  - `MIXED` status would indicate the different status of clone on different locales, and individual target locale statues should be checked.
  - `FAILED` status would indicate the source locale was not cloned all target locales successfully despite the request being eligible due to internal service issues.
  - `ROLLBACK_SUCCEEDED` status would indicate the skill was rolled back to the previous state in case any failure.
  - `ROLLBACK_FAILED` status would indicate that in case of failure, the rollback to the previous state of the skill was attempted, but it failed.
*/
type CloneLocaleRequestStatus string

func CloneLocaleRequestStatus_FAILED() CloneLocaleRequestStatus {
	return "FAILED"
}
func CloneLocaleRequestStatus_INELIGIBLE() CloneLocaleRequestStatus {
	return "INELIGIBLE"
}
func CloneLocaleRequestStatus_IN_PROGRESS() CloneLocaleRequestStatus {
	return "IN_PROGRESS"
}
func CloneLocaleRequestStatus_MIXED() CloneLocaleRequestStatus {
	return "MIXED"
}
func CloneLocaleRequestStatus_ROLLBACK_FAILED() CloneLocaleRequestStatus {
	return "ROLLBACK_FAILED"
}
func CloneLocaleRequestStatus_ROLLBACK_SUCCEEDED() CloneLocaleRequestStatus {
	return "ROLLBACK_SUCCEEDED"
}
func CloneLocaleRequestStatus_SUCCEEDED() CloneLocaleRequestStatus {
	return "SUCCEEDED"
}
