package catalog

type CreateContentUploadUrlRequest struct {
	// Provides the number of parts the file will be split into. An equal number of presigned upload urls are generated in response to facilitate each part's upload.
	NumberOfUploadParts int `json:"numberOfUploadParts"`
}

/*
{
 "properties": {
  "numberOfUploadParts": {
   "description": "Provides the number of parts the file will be split into. An equal number of presigned upload urls are generated in response to facilitate each part's upload.",
   "example": 1,
   "maximum": 10000,
   "minimum": 1,
   "type": "integer"
  }
 },
 "required": [
  "numberOfUploadParts"
 ],
 "type": "object"
}
*/
