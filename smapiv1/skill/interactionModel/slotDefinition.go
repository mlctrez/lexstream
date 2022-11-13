package interactionmodel

/*
SlotDefinition Slot definition.
*/
type SlotDefinition struct {
	// The name of the slot.
	Name string `json,omitempty:"name"`
	// The type of the slot. It can be a built-in or custom type.
	Type_ string `json,omitempty:"type"`
	// Configuration object for multiple values capturing behavior for this slot.
	MultipleValues *MultipleValuesConfig `json,omitempty:"multipleValues"`
	// Phrases the user can speak e.g. to trigger an intent or provide value for a slot elicitation.
	Samples []string `json,omitempty:"samples"`
}

/*
{
 "description": "Slot definition.",
 "properties": {
  "multipleValues": {
   "$ref": "#/definitions/v1.skill.interactionModel.MultipleValuesConfig",
   "description": "Configuration object for multiple values capturing behavior for this slot."
  },
  "name": {
   "description": "The name of the slot.",
   "type": "string"
  },
  "samples": {
   "description": "Phrases the user can speak e.g. to trigger an intent or provide value for a slot elicitation.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "type": {
   "description": "The type of the slot. It can be a built-in or custom type.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
