package typeversion

/*
SlotTypeUpdateObject Slot Type update description object.
*/
type SlotTypeUpdateObject struct {
	// The slot type description with a 255 character maximum.
	Description string `json,omitempty:"description"`
}

/*
{
 "description": "Slot Type update description object.",
 "properties": {
  "description": {
   "description": "The slot type description with a 255 character maximum.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
