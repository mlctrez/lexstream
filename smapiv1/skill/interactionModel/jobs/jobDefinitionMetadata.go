package jobs

/*
JobDefinitionMetadata Metadata of the job definition.
*/
type JobDefinitionMetadata struct {
	// Job identifier.
	Id string `json,omitempty:"id"`
	// Polymorphic type of the job.
	Type_  string               `json,omitempty:"type"`
	Status *JobDefinitionStatus `json,omitempty:"status"`
}

/*
{
 "description": "Metadata of the job definition.",
 "properties": {
  "id": {
   "description": "Job identifier.",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobDefinitionStatus",
   "x-isEnum": true
  },
  "type": {
   "description": "Polymorphic type of the job.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
