package annotationsets

type CreateNLUAnnotationSetRequest struct {
	// The locale of the NLU annotation set
	Locale string `json,omitempty:"locale"`
	// The name of NLU annotation set provided by customer
	Name string `json,omitempty:"name"`
}

/*
{
 "properties": {
  "locale": {
   "description": "The locale of the NLU annotation set",
   "type": "string"
  },
  "name": {
   "description": "The name of NLU annotation set provided by customer",
   "type": "string"
  }
 },
 "required": [
  "locale",
  "name"
 ],
 "type": "object"
}
*/
