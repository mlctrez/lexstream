package smarthomeevaluation

type PagedResponse struct {
	PaginationContext *PaginationContext `json:"paginationContext"`
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
