package jobs

/*
ExecutionMetadata ExecutionMetadata for executions.
*/
type ExecutionMetadata struct {
	// Identifier of the job.
	JobId string `json,omitempty:"jobId"`
	// ErrorCode to explain what went wrong in case of FAILUREs.
	ErrorCode string `json,omitempty:"errorCode"`
	// Current status of the job execution.
	Status string `json,omitempty:"status"`
}

/*
{
 "description": "ExecutionMetadata for executions.",
 "properties": {
  "errorCode": {
   "description": "ErrorCode to explain what went wrong in case of FAILUREs.",
   "type": "string"
  },
  "jobId": {
   "description": "Identifier of the job.",
   "type": "string"
  },
  "status": {
   "description": "Current status of the job execution.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
