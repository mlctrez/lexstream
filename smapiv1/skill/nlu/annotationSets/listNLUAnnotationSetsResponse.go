package annotationsets

type ListNLUAnnotationSetsResponse struct {
	AnnotationSets    []*AnnotationSet   `json:"annotationSets"`
	PaginationContext *PaginationContext `json:"paginationContext"`
	Links             *Links             `json:"_links"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.nlu.annotationSets.Links"
  },
  "annotationSets": {
   "items": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.AnnotationSet"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.nlu.annotationSets.PaginationContext"
  }
 },
 "type": "object"
}
*/
