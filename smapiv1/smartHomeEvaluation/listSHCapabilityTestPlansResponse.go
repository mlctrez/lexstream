package smarthomeevaluation

type ListSHCapabilityTestPlansResponse struct {
	PagedResponse
	TestPlans []*ListSHTestPlanItem `json:"testPlans,omitempty"`
}

/*
{
 "properties": {
  "testPlans": {
   "items": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.ListSHTestPlanItem"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
