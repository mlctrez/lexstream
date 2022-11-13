package smarthomeevaluation

import "time"

type GetSHCapabilityEvaluationResponse struct {
	EndTimestamp   time.Time                            `json:"endTimestamp,omitempty"`
	StartTimestamp time.Time                            `json:"startTimestamp,omitempty"`
	Status         *EvaluationEntityStatus              `json:"status,omitempty"`
	Error          *SHCapabilityError                   `json:"error,omitempty"`
	Input          *EvaluateSHCapabilityRequest         `json:"input,omitempty"`
	Metrics        map[string]SHEvaluationResultsMetric `json:"metrics,omitempty"`
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
