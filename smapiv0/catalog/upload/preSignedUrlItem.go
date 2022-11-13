package upload

type PreSignedUrlItem struct {
	ETag       string `json,omitempty:"eTag"`
	PartNumber int    `json,omitempty:"partNumber"`
}

/*
{
 "properties": {
  "eTag": {
   "type": "string"
  },
  "partNumber": {
   "type": "integer"
  }
 },
 "type": "object"
}
*/
