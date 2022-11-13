package annotationsets

type ListASRAnnotationSetsResponse struct {
	AnnotationSets    []*AnnotationSetItems `json:"annotationSets,omitempty"`
	PaginationContext *PaginationContext    `json:"paginationContext,omitempty"`
}

/*
{
 "properties": {
  "annotationSets": {
   "items": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.AnnotationSetItems"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.asr.annotationSets.PaginationContext"
  }
 },
 "required": [
  "annotationSets"
 ],
 "type": "object"
}
*/
