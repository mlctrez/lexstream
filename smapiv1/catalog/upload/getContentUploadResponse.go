package upload

import "time"

type GetContentUploadResponse struct {
	// Unique identifier of the upload
	Id string `json:"id"`
	// Unique identifier of the added catalog object
	CatalogId       string                    `json:"catalogId"`
	Status          *UploadStatus             `json:"status"`
	CreatedDate     time.Time                 `json:"createdDate"`
	LastUpdatedDate time.Time                 `json:"lastUpdatedDate"`
	File            *ContentUploadFileSummary `json:"file"`
	// List of different steps performed on the upload.
	IngestionSteps []*UploadIngestionStep `json:"ingestionSteps"`
}

/*
{
 "properties": {
  "catalogId": {
   "description": "Unique identifier of the added catalog object",
   "type": "string"
  },
  "createdDate": {
   "format": "date-time",
   "type": "string"
  },
  "file": {
   "$ref": "#/definitions/v1.catalog.upload.ContentUploadFileSummary"
  },
  "id": {
   "description": "Unique identifier of the upload",
   "type": "string"
  },
  "ingestionSteps": {
   "description": "List of different steps performed on the upload.",
   "items": {
    "$ref": "#/definitions/v1.catalog.upload.UploadIngestionStep"
   },
   "type": "array"
  },
  "lastUpdatedDate": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.catalog.upload.UploadStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "catalogId",
  "createdDate",
  "file",
  "id",
  "ingestionSteps",
  "lastUpdatedDate",
  "status"
 ],
 "type": "object"
}
*/
