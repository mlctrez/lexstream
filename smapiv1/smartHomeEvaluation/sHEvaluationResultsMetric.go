package smarthomeevaluation

type SHEvaluationResultsMetric struct {
	ErrorTestCases  int `json:"errorTestCases,omitempty"`
	FailedTestCases int `json:"failedTestCases,omitempty"`
	PassedTestCases int `json:"passedTestCases,omitempty"`
	TotalTestCases  int `json:"totalTestCases,omitempty"`
}

/*
{
 "properties": {
  "errorTestCases": {
   "type": "integer"
  },
  "failedTestCases": {
   "type": "integer"
  },
  "passedTestCases": {
   "type": "integer"
  },
  "totalTestCases": {
   "type": "integer"
  }
 },
 "type": "object"
}
*/
