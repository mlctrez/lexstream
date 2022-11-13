package smarthomeevaluation

type SHEvaluationResultsMetric struct {
	ErrorTestCases  int `json:"errorTestCases"`
	FailedTestCases int `json:"failedTestCases"`
	PassedTestCases int `json:"passedTestCases"`
	TotalTestCases  int `json:"totalTestCases"`
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
