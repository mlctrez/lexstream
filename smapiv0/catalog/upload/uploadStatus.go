package upload

/*
UploadStatus Status of the entire upload.
*/
type UploadStatus string

func UploadStatus_PENDING() UploadStatus {
	return "PENDING"
}
func UploadStatus_PROCESSING() UploadStatus {
	return "PROCESSING"
}
func UploadStatus_FAILED() UploadStatus {
	return "FAILED"
}
func UploadStatus_SUCCEEDED() UploadStatus {
	return "SUCCEEDED"
}
