package evaluations

type PagedResponse struct {
	PaginationContext *PaginationContext `json,omitempty:"paginationContext"`
	Links             *Links             `json,omitempty:"_links"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Links"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.PaginationContext"
  }
 },
 "type": "object"
}
*/
