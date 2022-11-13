package upload

/*
FileUploadStatus Value of status depends on if file is available for download or not.
*/
type FileUploadStatus string

func FileUploadStatus_PENDING() FileUploadStatus {
	return "PENDING"
}
func FileUploadStatus_AVAILABLE() FileUploadStatus {
	return "AVAILABLE"
}
func FileUploadStatus_PURGED() FileUploadStatus {
	return "PURGED"
}
func FileUploadStatus_UNAVAILABLE() FileUploadStatus {
	return "UNAVAILABLE"
}
