package evaluations

type PagedResultsResponse struct {
	PaginationContext *PagedResultsResponsePaginationContext `json:"paginationContext,omitempty"`
	Links             *Links                                 `json:"_links,omitempty"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Links"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.PagedResultsResponsePaginationContext"
  }
 },
 "type": "object"
}
*/
