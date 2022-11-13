package upload

/*
UploadStatus Status of the entire upload.
*/
type UploadStatus string

func UploadStatus_PENDING() UploadStatus {
	return "PENDING"
}
func UploadStatus_IN_PROGRESS() UploadStatus {
	return "IN_PROGRESS"
}
func UploadStatus_FAILED() UploadStatus {
	return "FAILED"
}
func UploadStatus_SUCCEEDED() UploadStatus {
	return "SUCCEEDED"
}
