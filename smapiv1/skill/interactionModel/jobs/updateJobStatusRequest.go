package jobs

/*
UpdateJobStatusRequest Update job status.
*/
type UpdateJobStatusRequest struct {
	Status *JobDefinitionStatus `json:"status,omitempty"`
}

/*
{
 "description": "Update job status.",
 "properties": {
  "status": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobDefinitionStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "status"
 ],
 "type": "object"
}
*/
