package skill

/*
CloneLocaleStatus Status for a locale clone on a particular target locale
  - `IN_PROGRESS` status would indicate the clone is still in progress to the target locale.
  - `SUCCEEDED` status would indicate the source locale was cloned successfully to the target locale.
  - `INELIGIBLE` status would indicate the source locale was ineligible to be cloned the target locale.
  - `FAILED` status would indicate the source locale was not cloned the target locale successfully.
  - `ROLLBACK_SUCCEEDED` status would indicate the locale was rolled back to the previous state in case any failure.
  - `ROLLBACK_FAILED` status would indicate that in case of failure, the rollback to the previous state of the locale was attempted, but it failed.
*/
type CloneLocaleStatus string

func CloneLocaleStatus_FAILED() CloneLocaleStatus {
	return "FAILED"
}
func CloneLocaleStatus_INELIGIBLE() CloneLocaleStatus {
	return "INELIGIBLE"
}
func CloneLocaleStatus_IN_PROGRESS() CloneLocaleStatus {
	return "IN_PROGRESS"
}
func CloneLocaleStatus_ROLLBACK_FAILED() CloneLocaleStatus {
	return "ROLLBACK_FAILED"
}
func CloneLocaleStatus_ROLLBACK_SUCCEEDED() CloneLocaleStatus {
	return "ROLLBACK_SUCCEEDED"
}
func CloneLocaleStatus_SUCCEEDED() CloneLocaleStatus {
	return "SUCCEEDED"
}
