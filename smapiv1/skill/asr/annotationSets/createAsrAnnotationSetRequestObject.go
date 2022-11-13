package annotationsets

type CreateAsrAnnotationSetRequestObject struct {
	// The name of ASR annotation set. The length of the name cannot exceed 170 chars. Only alphanumeric characters are accepted.
	Name string `json:"name"`
}

/*
{
 "properties": {
  "name": {
   "description": "The name of ASR annotation set. The length of the name cannot exceed 170 chars. Only alphanumeric characters are accepted.",
   "type": "string"
  }
 },
 "required": [
  "name"
 ],
 "type": "object"
}
*/
