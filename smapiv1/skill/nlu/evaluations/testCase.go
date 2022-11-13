package evaluations

type TestCase struct {
	Status   *ResultsStatus `json:"status"`
	Inputs   *Inputs        `json:"inputs"`
	Actual   *Actual        `json:"actual"`
	Expected []*Expected    `json:"expected"`
}

/*
{
 "properties": {
  "actual": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Actual"
  },
  "expected": {
   "items": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.Expected"
   },
   "type": "array"
  },
  "inputs": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Inputs"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.ResultsStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
