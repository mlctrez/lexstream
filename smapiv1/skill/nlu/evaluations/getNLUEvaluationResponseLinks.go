package evaluations

type GetNLUEvaluationResponseLinks struct {
	Results *Results `json:"results,omitempty"`
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
