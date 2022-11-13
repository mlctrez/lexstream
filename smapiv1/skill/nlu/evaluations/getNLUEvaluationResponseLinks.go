package evaluations

type GetNLUEvaluationResponseLinks struct {
	Results *Results `json,omitempty:"results"`
}

/*
{
 "properties": {
  "results": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Results"
  }
 },
 "type": "object"
}
*/
