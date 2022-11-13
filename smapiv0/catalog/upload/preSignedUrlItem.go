package upload

type PreSignedUrlItem struct {
	ETag       string `json:"eTag,omitempty"`
	PartNumber int    `json:"partNumber,omitempty"`
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
