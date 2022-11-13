package annotationsets

type ListASRAnnotationSetsResponse struct {
	AnnotationSets    []*AnnotationSetItems `json,omitempty:"annotationSets"`
	PaginationContext *PaginationContext    `json,omitempty:"paginationContext"`
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
