package jobs

/*
CreateJobDefinitionResponse The response of create job definition.
*/
type CreateJobDefinitionResponse struct {
	// Idenitifier of the job definition.
	JobId string `json:"jobId"`
}

/*
{
 "description": "The response of create job definition.",
 "properties": {
  "jobId": {
   "description": "Idenitifier of the job definition.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
