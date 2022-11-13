package smarthomeevaluation

type EvaluateSHCapabilityRequest struct {
	CapabilityTestPlan *CapabilityTestPlan `json,omitempty:"capabilityTestPlan"`
	Endpoint           *Endpoint           `json,omitempty:"endpoint"`
	Stage              *Stage              `json,omitempty:"stage"`
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
