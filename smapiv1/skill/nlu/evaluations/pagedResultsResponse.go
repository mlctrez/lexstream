package evaluations

type PagedResultsResponse struct {
	PaginationContext *PagedResultsResponsePaginationContext `json,omitempty:"paginationContext"`
	Links             *Links                                 `json,omitempty:"_links"`
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
