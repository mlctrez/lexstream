package jobs

/*
JobDefinitionStatus Current status of the job definition.
*/
type JobDefinitionStatus string

func JobDefinitionStatus_DISABLED() JobDefinitionStatus {
	return "DISABLED"
}
func JobDefinitionStatus_ENALBED() JobDefinitionStatus {
	return "ENALBED"
}
