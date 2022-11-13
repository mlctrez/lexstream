package smarthomeevaluation

type PagedResponse struct {
	PaginationContext *PaginationContext `json:"paginationContext,omitempty"`
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
