package upload

import "time"

type ContentUploadSummary struct {
	// Unique identifier of the upload.
	Id string `json:"id,omitempty"`
	// Provides a unique identifier of the catalog.
	CatalogId       string        `json:"catalogId,omitempty"`
	Status          *UploadStatus `json:"status,omitempty"`
	CreatedDate     time.Time     `json:"createdDate,omitempty"`
	LastUpdatedDate time.Time     `json:"lastUpdatedDate,omitempty"`
}

/*
{
 "properties": {
  "catalogId": {
   "description": "Provides a unique identifier of the catalog.",
   "type": "string"
  },
  "createdDate": {
   "format": "date-time",
   "type": "string"
  },
  "id": {
   "description": "Unique identifier of the upload.",
   "type": "string"
  },
  "lastUpdatedDate": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v0.catalog.upload.UploadStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
