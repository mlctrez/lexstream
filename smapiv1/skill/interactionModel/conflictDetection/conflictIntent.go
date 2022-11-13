package conflictdetection

type ConflictIntent struct {
	// Conflict intent name
	Name string `json:"name,omitempty"`
	// List of conflict intent slots
	Slots map[string]ConflictIntentSlot `json:"slots,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "Conflict intent name",
   "type": "string"
  },
  "slots": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.ConflictIntentSlot"
   },
   "description": "List of conflict intent slots",
   "type": "object"
  }
 },
 "required": [
  "name"
 ],
 "type": "object"
}
*/
