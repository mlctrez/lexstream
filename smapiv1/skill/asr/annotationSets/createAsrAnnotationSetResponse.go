package annotationsets

type CreateAsrAnnotationSetResponse struct {
	// ID used to retrieve the ASR annotation set.
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "ID used to retrieve the ASR annotation set.",
   "type": "string"
  }
 },
 "required": [
  "id"
 ],
 "type": "object"
}
*/
