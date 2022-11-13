package skill

/*
RollbackRequestStatusTypes The rollback status of the rollback request.
* `FAILED` - The rollback has failed. Please retry the rollback.
* `INELIGIBLE` - The target version is ineligible for rollback.
* `IN_PROGRESS` - The rollback is in progress.
* `SUCCEEDED` - The rollback has succeeded.
*/
type RollbackRequestStatusTypes string

func RollbackRequestStatusTypes_FAILED() RollbackRequestStatusTypes {
	return "FAILED"
}
func RollbackRequestStatusTypes_INELIGIBLE() RollbackRequestStatusTypes {
	return "INELIGIBLE"
}
func RollbackRequestStatusTypes_IN_PROGRESS() RollbackRequestStatusTypes {
	return "IN_PROGRESS"
}
func RollbackRequestStatusTypes_SUCCEEDED() RollbackRequestStatusTypes {
	return "SUCCEEDED"
}
