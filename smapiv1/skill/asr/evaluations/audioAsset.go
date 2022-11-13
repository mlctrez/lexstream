package evaluations

/*
AudioAsset Object containing information about downloading audio file
*/
type AudioAsset struct {
	// S3 presigned download url for downloading the audio file
	DownloadUrl string `json:"downloadUrl"`
	// timestamp when the audio download url expire in ISO 8601 format
	ExpiryTime string `json:"expiryTime"`
}

/*
{
 "description": "Object containing information about downloading audio file",
 "properties": {
  "downloadUrl": {
   "description": "S3 presigned download url for downloading the audio file",
   "type": "string"
  },
  "expiryTime": {
   "description": "timestamp when the audio download url expire in ISO 8601 format",
   "type": "string"
  }
 },
 "required": [
  "downloadUrl",
  "expiryTime"
 ],
 "type": "object"
}
*/