package smarthomeevaluation

type EvaluateSHCapabilityRequest struct {
	CapabilityTestPlan *CapabilityTestPlan `json:"capabilityTestPlan,omitempty"`
	Endpoint           *Endpoint           `json:"endpoint,omitempty"`
	Stage              *Stage              `json:"stage,omitempty"`
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
