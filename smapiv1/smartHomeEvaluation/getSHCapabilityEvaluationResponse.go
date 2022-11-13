package smarthomeevaluation

import "time"

type GetSHCapabilityEvaluationResponse struct {
	EndTimestamp   time.Time                            `json:"endTimestamp"`
	StartTimestamp time.Time                            `json:"startTimestamp"`
	Status         *EvaluationEntityStatus              `json:"status"`
	Error          *SHCapabilityError                   `json:"error"`
	Input          *EvaluateSHCapabilityRequest         `json:"input"`
	Metrics        map[string]SHEvaluationResultsMetric `json:"metrics"`
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
