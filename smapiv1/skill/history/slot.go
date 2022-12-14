package history

type Slot struct {
	// Name of the slot that was used in this interaction.
	Name string `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "Name of the slot that was used in this interaction.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
