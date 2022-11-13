package jobs

/*
JobDefinition Definition for dynamic job.
*/
type JobDefinition struct {
	// Polymorphic type of the job
	Type_   string   `json:"type,omitempty"`
	Trigger *Trigger `json:"trigger,omitempty"`
	// Current status of the job definition.
	Status string `json:"status,omitempty"`
}

/*
{
 "description": "Definition for dynamic job.",
 "discriminator": "type",
 "properties": {
  "status": {
   "description": "Current status of the job definition.",
   "type": "string"
  },
  "trigger": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.Trigger"
  },
  "type": {
   "description": "Polymorphic type of the job",
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "type": "object"
}
*/
