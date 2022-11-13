package upload

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
UploadIngestionStep Represents a single step in the multi-step ingestion process of a new upload.
*/
type UploadIngestionStep struct {
	Name   *IngestionStepName `json:"name,omitempty"`
	Status *IngestionStatus   `json:"status,omitempty"`
	// Url for the file containing logs of ingestion step.
	LogUrl string `json:"logUrl,omitempty"`
	// This array will contain the violations occurred during the execution of step. Will be empty, if execution succeeded.
	Violations []*smapiv1.Error `json:"violations,omitempty"`
}

/*
{
 "description": "Represents a single step in the multi-step ingestion process of a new upload.",
 "properties": {
  "logUrl": {
   "description": "Url for the file containing logs of ingestion step.",
   "format": "uri",
   "type": "string"
  },
  "name": {
   "$ref": "#/definitions/v1.catalog.upload.IngestionStepName",
   "x-isEnum": true
  },
  "status": {
   "$ref": "#/definitions/v1.catalog.upload.IngestionStatus",
   "x-isEnum": true
  },
  "violations": {
   "description": "This array will contain the violations occurred during the execution of step. Will be empty, if execution succeeded.",
   "items": {
    "$ref": "#/definitions/v1.Error"
   },
   "type": "array"
  }
 },
 "required": [
  "name",
  "status",
  "violations"
 ],
 "type": "object"
}
*/
