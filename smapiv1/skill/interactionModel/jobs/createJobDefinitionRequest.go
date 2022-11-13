package jobs

/*
CreateJobDefinitionRequest Request to create job definitions.
*/
type CreateJobDefinitionRequest struct {
	// ID of the vendor owning the skill.
	VendorId      string         `json,omitempty:"vendorId"`
	JobDefinition *JobDefinition `json,omitempty:"jobDefinition"`
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
