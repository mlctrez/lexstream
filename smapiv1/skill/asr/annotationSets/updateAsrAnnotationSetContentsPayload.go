package annotationsets

/*
UpdateAsrAnnotationSetContentsPayload This is the payload shema for updating asr annotation set contents. Note for text/csv content type, the  csv header definitions need to follow the properties of '#/definitions/Annotaion'
*/
type UpdateAsrAnnotationSetContentsPayload struct {
	Annotations []*Annotation `json:"annotations,omitempty"`
}

/*
{
 "description": "This is the payload shema for updating asr annotation set contents. Note for text/csv content type, the  csv header definitions need to follow the properties of '#/definitions/Annotaion'\n",
 "properties": {
  "annotations": {
   "items": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.Annotation"
   },
   "type": "array"
  }
 },
 "required": [
  "annotations"
 ],
 "type": "object"
}
*/
