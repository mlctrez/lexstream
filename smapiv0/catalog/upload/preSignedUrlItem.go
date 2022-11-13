package upload

type PreSignedUrlItem struct {
	ETag       string `json:"eTag"`
	PartNumber int    `json:"partNumber"`
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
