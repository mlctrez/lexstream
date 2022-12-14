package interactionmodel

/*
Intent The set of intents your service can accept and process.
*/
type Intent struct {
	// Name to identify the intent.
	Name string `json:"name,omitempty"`
	// List of slots within the intent.
	Slots []*SlotDefinition `json:"slots,omitempty"`
	// Phrases the user can speak e.g. to trigger an intent or provide value for a slot elicitation.
	Samples []string `json:"samples,omitempty"`
}

/*
{
 "description": "The set of intents your service can accept and process.",
 "properties": {
  "name": {
   "description": "Name to identify the intent.",
   "type": "string"
  },
  "samples": {
   "description": "Phrases the user can speak e.g. to trigger an intent or provide value for a slot elicitation.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "slots": {
   "description": "List of slots within the intent.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.SlotDefinition"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
