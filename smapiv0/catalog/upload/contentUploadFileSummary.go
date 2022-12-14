package upload

type ContentUploadFileSummary struct {
	// If the file is available for download, presigned download URL can be used to download the file.
	PresignedDownloadUrl string            `json:"presignedDownloadUrl,omitempty"`
	Status               *FileUploadStatus `json:"status,omitempty"`
}

/*
{
 "properties": {
  "presignedDownloadUrl": {
   "description": "If the file is available for download, presigned download URL can be used to download the file.",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v0.catalog.upload.FileUploadStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
