package upload

type IngestionStatus string

func IngestionStatus_PENDING() IngestionStatus {
	return "PENDING"
}
func IngestionStatus_IN_PROGRESS() IngestionStatus {
	return "IN_PROGRESS"
}
func IngestionStatus_FAILED() IngestionStatus {
	return "FAILED"
}
func IngestionStatus_SUCCEEDED() IngestionStatus {
	return "SUCCEEDED"
}
func IngestionStatus_CANCELLED() IngestionStatus {
	return "CANCELLED"
}
