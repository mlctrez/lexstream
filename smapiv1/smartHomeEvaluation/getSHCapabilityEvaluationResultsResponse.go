package smarthomeevaluation

type GetSHCapabilityEvaluationResultsResponse struct {
	PagedResponse
	TestCaseResults []*TestCaseResult `json:"testCaseResults,omitempty"`
}

/*
{
 "properties": {
  "testCaseResults": {
   "items": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.TestCaseResult"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
