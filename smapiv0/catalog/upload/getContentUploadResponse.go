package upload

/*
GetContentUploadResponse Response object for get content upload request.
*/
type GetContentUploadResponse struct {
	ContentUploadSummary
	File           *ContentUploadFileSummary `json:"file"`
	IngestionSteps []*UploadIngestionStep    `json:"ingestionSteps"`
}

/*
{
 "properties": {
  "file": {
   "$ref": "#/definitions/v0.catalog.upload.ContentUploadFileSummary"
  },
  "ingestionSteps": {
   "items": {
    "$ref": "#/definitions/v0.catalog.upload.UploadIngestionStep"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
