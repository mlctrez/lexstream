package resourceschema

import "time"

type GetResourceSchemaResponse struct {
	// S3 presigned URL to schema location.
	SchemaLocationUrl string `json,omitempty:"schemaLocationUrl"`
	// Timestamp when the schema location url expires in ISO 8601 format.
	ExpiryTime time.Time `json,omitempty:"expiryTime"`
}

/*
{
 "properties": {
  "expiryTime": {
   "description": "Timestamp when the schema location url expires in ISO 8601 format.",
   "example": "2020-10-19T23:50:02.135Z",
   "format": "date-time",
   "type": "string"
  },
  "schemaLocationUrl": {
   "description": "S3 presigned URL to schema location.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
