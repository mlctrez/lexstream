package upload

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

/*
UploadIngestionStep Represents a single step in the ingestion process of a new upload.
*/
type UploadIngestionStep struct {
	Name   *IngestionStepName `json:"name"`
	Status *IngestionStatus   `json:"status"`
	// Represents the url for the file containing logs of ingestion step.
	LogUrl string `json:"logUrl"`
	// This array will contain the errors occurred during the execution of step. Will be empty, if execution succeeded.
	Errors []*smapiv0.Error `json:"errors"`
}

/*
{
 "description": "Represents a single step in the ingestion process of a new upload.",
 "properties": {
  "errors": {
   "description": "This array will contain the errors occurred during the execution of step. Will be empty, if execution succeeded.",
   "items": {
    "$ref": "#/definitions/v0.Error"
   },
   "type": "array"
  },
  "logUrl": {
   "description": "Represents the url for the file containing logs of ingestion step.",
   "type": "string"
  },
  "name": {
   "$ref": "#/definitions/v0.catalog.upload.IngestionStepName",
   "x-isEnum": true
  },
  "status": {
   "$ref": "#/definitions/v0.catalog.upload.IngestionStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "errors",
  "name",
  "status"
 ],
 "type": "object"
}
*/
