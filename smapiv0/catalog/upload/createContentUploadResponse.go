package upload

/*
CreateContentUploadResponse Request body for self-hosted catalog uploads.
*/
type CreateContentUploadResponse struct {
	ContentUploadSummary
	IngestionSteps []*UploadIngestionStep `json,omitempty:"ingestionSteps"`
	// Ordered list of presigned upload parts to perform a partitioned (multipart) file upload.
	PresignedUploadParts []*PresignedUploadPart `json,omitempty:"presignedUploadParts"`
}

/*
{
 "properties": {
  "ingestionSteps": {
   "items": {
    "$ref": "#/definitions/v0.catalog.upload.UploadIngestionStep"
   },
   "type": "array"
  },
  "presignedUploadParts": {
   "description": "Ordered list of presigned upload parts to perform a partitioned (multipart) file upload.",
   "items": {
    "$ref": "#/definitions/v0.catalog.upload.PresignedUploadPart"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
