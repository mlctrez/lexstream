package smarthomeevaluation

import "time"

type EvaluationObject struct {
	EndTimestamp      time.Time               `json:"endTimestamp"`
	StartTimestamp    time.Time               `json:"startTimestamp"`
	Status            *EvaluationEntityStatus `json:"status"`
	EndpointId        string                  `json:"endpointId"`
	Id                string                  `json:"id"`
	SourceTestPlanIds []string                `json:"sourceTestPlanIds"`
	TestPlanId        string                  `json:"testPlanId"`
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
