package upload

/*
PresignedUploadPart Single upload part to perform a partitioned (multipart) file upload.
*/
type PresignedUploadPart struct {
	Url        string `json:"url"`
	PartNumber int    `json:"partNumber"`
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
