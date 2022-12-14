package annotationsets

type UpdateNLUAnnotationSetPropertiesRequest struct {
	// The name of NLU annotation set provided by customer
	Name string `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "The name of NLU annotation set provided by customer",
   "type": "string"
  }
 },
 "required": [
  "name"
 ],
 "type": "object"
}
*/
