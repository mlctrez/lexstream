package conflictdetection

/*
ConflictDetectionJobStatus The status of conflict detection job.
*/
type ConflictDetectionJobStatus string

func ConflictDetectionJobStatus_IN_PROGRESS() ConflictDetectionJobStatus {
	return "IN_PROGRESS"
}
func ConflictDetectionJobStatus_COMPLETED() ConflictDetectionJobStatus {
	return "COMPLETED"
}
func ConflictDetectionJobStatus_FAILED() ConflictDetectionJobStatus {
	return "FAILED"
}
