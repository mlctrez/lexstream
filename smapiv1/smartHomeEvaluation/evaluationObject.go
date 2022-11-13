package smarthomeevaluation

import "time"

type EvaluationObject struct {
	EndTimestamp      time.Time               `json,omitempty:"endTimestamp"`
	StartTimestamp    time.Time               `json,omitempty:"startTimestamp"`
	Status            *EvaluationEntityStatus `json,omitempty:"status"`
	EndpointId        string                  `json,omitempty:"endpointId"`
	Id                string                  `json,omitempty:"id"`
	SourceTestPlanIds []string                `json,omitempty:"sourceTestPlanIds"`
	TestPlanId        string                  `json,omitempty:"testPlanId"`
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
