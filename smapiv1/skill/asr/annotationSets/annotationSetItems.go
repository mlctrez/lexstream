package annotationsets

type AnnotationSetItems struct {
	AnnotationSetMetadata
	// The Annotation set id
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "The Annotation set id",
   "type": "string"
  }
 },
 "required": [
  "id"
 ],
 "type": "object"
}
*/
