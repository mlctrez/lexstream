package upload

import "time"

type ContentUploadSummary struct {
	// Unique identifier of the upload.
	Id string `json,omitempty:"id"`
	// Provides a unique identifier of the catalog.
	CatalogId       string        `json,omitempty:"catalogId"`
	Status          *UploadStatus `json,omitempty:"status"`
	CreatedDate     time.Time     `json,omitempty:"createdDate"`
	LastUpdatedDate time.Time     `json,omitempty:"lastUpdatedDate"`
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
