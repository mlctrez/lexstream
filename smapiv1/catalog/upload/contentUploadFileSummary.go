package upload

import "time"

type ContentUploadFileSummary struct {
	// If the file is available for download, downloadUrl can be used to download the file.
	DownloadUrl string            `json:"downloadUrl,omitempty"`
	ExpiresAt   time.Time         `json:"expiresAt,omitempty"`
	Status      *FileUploadStatus `json:"status,omitempty"`
}

/*
{
 "properties": {
  "downloadUrl": {
   "description": "If the file is available for download, downloadUrl can be used to download the file.",
   "format": "uri",
   "type": "string"
  },
  "expiresAt": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.catalog.upload.FileUploadStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "downloadUrl",
  "expiresAt",
  "status"
 ],
 "type": "object"
}
*/
