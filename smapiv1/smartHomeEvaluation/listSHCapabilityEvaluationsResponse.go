package smarthomeevaluation

type ListSHCapabilityEvaluationsResponse struct {
	PaginationContextToken *PaginationContextToken `json,omitempty:"paginationContextToken"`
	Evaluations            []*EvaluationObject     `json,omitempty:"evaluations"`
}

/*
{
 "properties": {
  "evaluations": {
   "items": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluationObject"
   },
   "type": "array"
  },
  "paginationContextToken": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.paginationContextToken"
  }
 },
 "type": "object"
}
*/
