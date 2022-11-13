package annotationsets

type ListNLUAnnotationSetsResponse struct {
	AnnotationSets    []*AnnotationSet   `json:"annotationSets,omitempty"`
	PaginationContext *PaginationContext `json:"paginationContext,omitempty"`
	Links             *Links             `json:"_links,omitempty"`
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
