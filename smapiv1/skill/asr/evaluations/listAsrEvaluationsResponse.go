package evaluations

/*
ListAsrEvaluationsResponse response body for a list evaluation API
*/
type ListAsrEvaluationsResponse struct {
	// an array containing all evaluations that have ever run by developers based on the filter criteria defined in the request
	Evaluations       []*EvaluationItems `json,omitempty:"evaluations"`
	PaginationContext *PaginationContext `json,omitempty:"paginationContext"`
}

/*
{
 "description": "response body for a list evaluation API",
 "properties": {
  "evaluations": {
   "description": "an array containing all evaluations that have ever run by developers based on the filter criteria defined in the request",
   "items": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationItems"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.PaginationContext"
  }
 },
 "required": [
  "evaluations"
 ],
 "type": "object"
}
*/
