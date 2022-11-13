package upload

type CompleteUploadRequest struct {
	// List of (eTag, part number) pairs for each part of the file uploaded.
	PartETags []*PreSignedUrlItem `json:"partETags"`
}

/*
{
 "properties": {
  "partETags": {
   "description": "List of (eTag, part number) pairs for each part of the file uploaded.",
   "items": {
    "$ref": "#/definitions/v0.catalog.upload.PreSignedUrlItem"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
