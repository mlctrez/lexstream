package evaluations

type GetNLUEvaluationResponse struct {
	EvaluationEntity
	Links *GetNLUEvaluationResponseLinks `json:"_links,omitempty"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.GetNLUEvaluationResponseLinks"
  }
 },
 "type": "object"
}
*/
