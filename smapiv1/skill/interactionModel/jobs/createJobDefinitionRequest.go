package jobs

/*
CreateJobDefinitionRequest Request to create job definitions.
*/
type CreateJobDefinitionRequest struct {
	// ID of the vendor owning the skill.
	VendorId      string         `json:"vendorId,omitempty"`
	JobDefinition *JobDefinition `json:"jobDefinition,omitempty"`
}

/*
{
 "description": "Request to create job definitions.",
 "properties": {
  "jobDefinition": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobDefinition"
  },
  "vendorId": {
   "description": "ID of the vendor owning the skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
