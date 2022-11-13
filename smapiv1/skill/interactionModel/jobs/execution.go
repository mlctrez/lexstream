package jobs

import "time"

/*
Execution Execution data.
*/
type Execution struct {
	// Identifier of the execution.
	ExecutionId string `json:"executionId,omitempty"`
	// ISO date-time timestamp when the execution starts.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// ErrorCode to explain what went wrong in case of FAILUREs.
	ErrorCode string `json:"errorCode,omitempty"`
	// Current status of the job execution.
	Status       string           `json:"status,omitempty"`
	ErrorDetails *JobErrorDetails `json:"errorDetails,omitempty"`
}

/*
{
 "description": "Execution data.",
 "properties": {
  "errorCode": {
   "description": "ErrorCode to explain what went wrong in case of FAILUREs.",
   "type": "string"
  },
  "errorDetails": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobErrorDetails"
  },
  "executionId": {
   "description": "Identifier of the execution.",
   "type": "string"
  },
  "status": {
   "description": "Current status of the job execution.",
   "type": "string"
  },
  "timestamp": {
   "description": "ISO date-time timestamp when the execution starts.",
   "example": "2020-10-25T20:00:02.135Z",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
