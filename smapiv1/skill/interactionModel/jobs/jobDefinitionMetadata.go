package jobs

/*
JobDefinitionMetadata Metadata of the job definition.
*/
type JobDefinitionMetadata struct {
	// Job identifier.
	Id string `json:"id,omitempty"`
	// Polymorphic type of the job.
	Type_  string               `json:"type,omitempty"`
	Status *JobDefinitionStatus `json:"status,omitempty"`
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
