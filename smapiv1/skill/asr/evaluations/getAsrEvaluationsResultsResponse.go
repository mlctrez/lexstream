package evaluations

/*
GetAsrEvaluationsResultsResponse response for GetAsrEvaluationsResults
*/
type GetAsrEvaluationsResultsResponse struct {
	// array containing all evaluation results.
	Results           []*EvaluationResult `json:"results,omitempty"`
	PaginationContext *PaginationContext  `json:"paginationContext,omitempty"`
}

/*
{
 "description": "response for GetAsrEvaluationsResults",
 "properties": {
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.PaginationContext"
  },
  "results": {
   "description": "array containing all evaluation results.",
   "items": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationResult"
   },
   "type": "array"
  }
 },
 "required": [
  "results"
 ],
 "type": "object"
}
*/
