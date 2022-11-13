package evaluations

/*
ResolutionsPerAuthorityValueItems An object representing the resolved value for the slot, based on the user's utterance and the slot type definition.
*/
type ResolutionsPerAuthorityValueItems struct {
	// The string for the resolved slot value.
	Name string `json:"name"`
	/*
	   The unique ID defined for the resolved slot value. This is based on the IDs defined in the slot type definition.
	*/
	Id string `json:"id"`
}

/*
{
 "description": "An object representing the resolved value for the slot, based on the user's utterance and the slot type definition.\n",
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
