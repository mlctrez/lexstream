package evaluations

import "time"

type EvaluationEntity struct {
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	Status         *Status   `json:"status"`
	// Error message when evaluation job fails
	ErrorMessage string            `json:"errorMessage"`
	Inputs       *EvaluationInputs `json:"inputs"`
}

/*
{
 "properties": {
  "endTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "errorMessage": {
   "description": "Error message when evaluation job fails",
   "type": "string"
  },
  "inputs": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.EvaluationInputs"
  },
  "startTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Status",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
