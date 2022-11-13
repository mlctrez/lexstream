package smarthomeevaluation

import "time"

type GetSHCapabilityEvaluationResponse struct {
	EndTimestamp   time.Time                            `json,omitempty:"endTimestamp"`
	StartTimestamp time.Time                            `json,omitempty:"startTimestamp"`
	Status         *EvaluationEntityStatus              `json,omitempty:"status"`
	Error          *SHCapabilityError                   `json,omitempty:"error"`
	Input          *EvaluateSHCapabilityRequest         `json,omitempty:"input"`
	Metrics        map[string]SHEvaluationResultsMetric `json,omitempty:"metrics"`
}

/*
{
 "properties": {
  "endTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "error": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityError"
  },
  "input": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluateSHCapabilityRequest"
  },
  "metrics": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.SHEvaluationResultsMetric"
   },
   "type": "object"
  },
  "startTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluationEntityStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
