package smarthomeevaluation

type PagedResponse struct {
	PaginationContext *PaginationContext `json,omitempty:"paginationContext"`
}

/*
{
 "properties": {
  "paginationContext": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.paginationContext"
  }
 },
 "type": "object"
}
*/
