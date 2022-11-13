package smarthomeevaluation

import "time"

type EvaluationObject struct {
	EndTimestamp      time.Time               `json:"endTimestamp,omitempty"`
	StartTimestamp    time.Time               `json:"startTimestamp,omitempty"`
	Status            *EvaluationEntityStatus `json:"status,omitempty"`
	EndpointId        string                  `json:"endpointId,omitempty"`
	Id                string                  `json:"id,omitempty"`
	SourceTestPlanIds []string                `json:"sourceTestPlanIds,omitempty"`
	TestPlanId        string                  `json:"testPlanId,omitempty"`
}

/*
{
 "properties": {
  "endTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "endpointId": {
   "type": "string"
  },
  "id": {
   "type": "string"
  },
  "sourceTestPlanIds": {
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "startTimestamp": {
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluationEntityStatus",
   "x-isEnum": true
  },
  "testPlanId": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
