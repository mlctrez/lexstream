package upload

/*
PresignedUploadPart Single upload part to perform a partitioned (multipart) file upload.
*/
type PresignedUploadPart struct {
	Url        string `json:"url,omitempty"`
	PartNumber int    `json:"partNumber,omitempty"`
}

/*
{
 "description": "Single upload part to perform a partitioned (multipart) file upload.",
 "properties": {
  "partNumber": {
   "type": "integer"
  },
  "url": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
