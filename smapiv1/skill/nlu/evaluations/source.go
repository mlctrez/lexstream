package evaluations

/*
Source Use Annotation Set as evaluation source
*/
type Source struct {
	// ID of the annotation set
	AnnotationId string `json,omitempty:"annotationId"`
}

/*
{
 "description": "Use Annotation Set as evaluation source\n",
 "properties": {
  "annotationId": {
   "description": "ID of the annotation set",
   "type": "string"
  }
 },
 "type": "object"
}
*/
