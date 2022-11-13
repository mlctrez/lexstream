package evaluations

import "time"

/*
EvaluationMetadata response body for GetAsrEvaluationsStatus API
*/
type EvaluationMetadata struct {
	Status *EvaluationStatus `json,omitempty:"status"`
	// indicate the total number of evaluations that are supposed to be run in the evaluation request
	TotalEvaluationCount int `json,omitempty:"totalEvaluationCount"`
	// indicate the number of completed evaluations
	CompletedEvaluationCount int `json,omitempty:"completedEvaluationCount"`
	// indicate the start time stamp of the ASR evaluation job. ISO-8601 Format.
	StartTimestamp time.Time                        `json,omitempty:"startTimestamp"`
	Request        *PostAsrEvaluationsRequestObject `json,omitempty:"request"`
	Error          *ErrorObject                     `json,omitempty:"error"`
	Result         *EvaluationMetadataResult        `json,omitempty:"result"`
}

/*
{
 "description": "response body for GetAsrEvaluationsStatus API",
 "properties": {
  "completedEvaluationCount": {
   "description": "indicate the number of completed evaluations",
   "type": "number"
  },
  "error": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.ErrorObject"
  },
  "request": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.PostAsrEvaluationsRequestObject"
  },
  "result": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationMetadataResult"
  },
  "startTimestamp": {
   "description": "indicate the start time stamp of the ASR evaluation job. ISO-8601 Format.",
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationStatus",
   "x-isEnum": true
  },
  "totalEvaluationCount": {
   "description": "indicate the total number of evaluations that are supposed to be run in the evaluation request",
   "type": "number"
  }
 },
 "required": [
  "completedEvaluationCount",
  "request",
  "startTimestamp",
  "status",
  "totalEvaluationCount"
 ],
 "type": "object"
}
*/
