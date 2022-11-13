package evaluations

type GetNLUEvaluationResultsResponse struct {
	PagedResultsResponse
	/*
	   count of tests failed. A test fails when the expected intent and expected slots are not identical.
	*/
	TotalFailed int         `json:"totalFailed,omitempty"`
	TestCases   []*TestCase `json:"testCases,omitempty"`
}

/*
{
 "properties": {
  "testCases": {
   "items": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.TestCase"
   },
   "type": "array"
  },
  "totalFailed": {
   "description": "count of tests failed. A test fails when the expected intent and expected slots are not identical.\n",
   "type": "number"
  }
 },
 "type": "object"
}
*/
