package smarthomeevaluation

type ListSHCapabilityTestPlansResponse struct {
	PagedResponse
	TestPlans []*ListSHTestPlanItem `json,omitempty:"testPlans"`
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
