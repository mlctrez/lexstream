package type1

/*
SlotTypeInput Definition for slot type input.
*/
type SlotTypeInput struct {
	// Name of the slot type.
	Name string `json:"name"`
	// Description string about the slot type.
	Description string `json:"description"`
}

/*
{
 "description": "Definition for slot type input.",
 "properties": {
  "description": {
   "description": "Description string about the slot type.",
   "type": "string"
  },
  "name": {
   "description": "Name of the slot type.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
