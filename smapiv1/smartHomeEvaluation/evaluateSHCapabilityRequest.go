package smarthomeevaluation

type EvaluateSHCapabilityRequest struct {
	CapabilityTestPlan *CapabilityTestPlan `json:"capabilityTestPlan"`
	Endpoint           *Endpoint           `json:"endpoint"`
	Stage              *Stage              `json:"stage"`
}

/*
{
 "properties": {
  "capabilityTestPlan": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.CapabilityTestPlan"
  },
  "endpoint": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.Endpoint"
  },
  "stage": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.Stage",
   "x-isEnum": true
  }
 },
 "required": [
  "capabilityTestPlan",
  "endpoint",
  "stage"
 ],
 "type": "object"
}
*/
