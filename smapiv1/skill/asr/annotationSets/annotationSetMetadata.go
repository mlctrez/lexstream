package annotationsets

import "time"

type AnnotationSetMetadata struct {
	// Name of the ASR annotation set
	Name string `json:"name,omitempty"`
	// Number of annotations within an annotation set
	AnnotationCount int `json:"annotationCount,omitempty"`
	// The timestamp for the most recent update of ASR annotation set
	LastUpdatedTimestamp time.Time `json:"lastUpdatedTimestamp,omitempty"`
	// Indicates if the annotation set is eligible for evaluation. A set is not eligible for evaluation if any annotation within the set has a missing uploadId, filePathInUpload, expectedTranscription, or evaluationWeight.
	EligibleForEvaluation bool `json:"eligibleForEvaluation,omitempty"`
}

/*
{
 "properties": {
  "annotationCount": {
   "description": "Number of annotations within an annotation set",
   "type": "integer"
  },
  "eligibleForEvaluation": {
   "description": "Indicates if the annotation set is eligible for evaluation. A set is not eligible for evaluation if any annotation within the set has a missing uploadId, filePathInUpload, expectedTranscription, or evaluationWeight.",
   "type": "boolean"
  },
  "lastUpdatedTimestamp": {
   "description": "The timestamp for the most recent update of ASR annotation set",
   "format": "date-time",
   "type": "string"
  },
  "name": {
   "description": "Name of the ASR annotation set",
   "type": "string"
  }
 },
 "required": [
  "annotationCount",
  "lastUpdatedTimestamp",
  "name"
 ],
 "type": "object"
}
*/
