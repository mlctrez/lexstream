package jobs

/*
JobErrorDetails Optional details if the execution is depending on other executions.
*/
type JobErrorDetails struct {
	ExecutionMetadata []*ExecutionMetadata `json:"executionMetadata"`
}

/*
{
 "description": "Optional details if the execution is depending on other executions.",
 "properties": {
  "executionMetadata": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.ExecutionMetadata"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
