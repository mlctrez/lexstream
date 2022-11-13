package evaluations

type TestCase struct {
	Status   *ResultsStatus `json,omitempty:"status"`
	Inputs   *Inputs        `json,omitempty:"inputs"`
	Actual   *Actual        `json,omitempty:"actual"`
	Expected []*Expected    `json,omitempty:"expected"`
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
