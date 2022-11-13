package catalog

type CreateContentUploadUrlResponse struct {
	// Unique identifier for collection of generated urls.
	UrlId string `json:"urlId,omitempty"`
	// Ordered list of presigned upload parts to perform a partitioned (multipart) file upload
	PreSignedUploadParts []*PresignedUploadPartItems `json:"preSignedUploadParts,omitempty"`
}

/*
{
 "properties": {
  "preSignedUploadParts": {
   "description": "Ordered list of presigned upload parts to perform a partitioned (multipart) file upload",
   "items": {
    "$ref": "#/definitions/v1.catalog.PresignedUploadPartItems"
   },
   "type": "array"
  },
  "urlId": {
   "description": "Unique identifier for collection of generated urls.",
   "type": "string"
  }
 },
 "required": [
  "urlId"
 ],
 "type": "object"
}
*/
