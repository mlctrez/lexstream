package annotationsets

type AnnotationSet struct {
	AnnotationSetEntity
	// Identifier of the NLU annotation set.
	AnnotationId string `json:"annotationId"`
}

/*
{
 "properties": {
  "annotationId": {
   "description": "Identifier of the NLU annotation set.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
