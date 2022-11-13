package type1

/*
SlotTypeUpdateDefinition Slot type update definition object.
*/
type SlotTypeUpdateDefinition struct {
	// The slot type description with a 255 character maximum.
	Description string `json:"description,omitempty"`
}

/*
{
 "description": "Slot type update definition object.",
 "properties": {
  "description": {
   "description": "The slot type description with a 255 character maximum.",
   "type": "string"
  }
 },
 "required": [
  "description"
 ],
 "type": "object"
}
*/
