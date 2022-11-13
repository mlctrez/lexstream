package evaluations

type ResolutionsPerAuthorityValue struct {
	// The string for the resolved slot value.
	Name string `json:"name,omitempty"`
	/*
	   The unique ID defined for the resolved slot value. This is based on the IDs defined in the slot type definition.
	*/
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "The unique ID defined for the resolved slot value. This is based on the IDs defined in the slot type definition.\n",
   "type": "string"
  },
  "name": {
   "description": "The string for the resolved slot value.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
