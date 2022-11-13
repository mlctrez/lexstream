package smarthomeevaluation

type SHEvaluationResultsMetric struct {
	ErrorTestCases  int `json,omitempty:"errorTestCases"`
	FailedTestCases int `json,omitempty:"failedTestCases"`
	PassedTestCases int `json,omitempty:"passedTestCases"`
	TotalTestCases  int `json,omitempty:"totalTestCases"`
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
