package evaluations

/*
ListNLUEvaluationsResponse response body for a list evaluation API
*/
type ListNLUEvaluationsResponse struct {
	PagedResponse
	Evaluations []*Evaluation `json,omitempty:"evaluations"`
}

/*
{
 "properties": {
  "evaluations": {
   "items": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.Evaluation"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
