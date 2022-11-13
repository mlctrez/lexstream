package interactionmodel

type DialogSlotItems struct {
	// The name of the slot that has dialog rules associated with it.
	Name string `json:"name"`
	// Type of the slot in the dialog intent.
	Type_ string `json:"type"`
	// Describes whether elicitation of the slot is required.
	ElicitationRequired bool `json:"elicitationRequired"`
	// Describes whether confirmation of the slot is required.
	ConfirmationRequired bool           `json:"confirmationRequired"`
	Prompts              *DialogPrompts `json:"prompts"`
	// List of validations for the slot. if validation fails, user will be prompted with the provided prompt.
	Validations []*SlotValidation `json:"validations"`
}

/*
{
 "properties": {
  "confirmationRequired": {
   "description": "Describes whether confirmation of the slot is required.",
   "type": "boolean"
  },
  "elicitationRequired": {
   "description": "Describes whether elicitation of the slot is required.",
   "type": "boolean"
  },
  "name": {
   "description": "The name of the slot that has dialog rules associated with it.",
   "type": "string"
  },
  "prompts": {
   "$ref": "#/definitions/v1.skill.interactionModel.DialogPrompts"
  },
  "type": {
   "description": "Type of the slot in the dialog intent.",
   "type": "string"
  },
  "validations": {
   "description": "List of validations for the slot. if validation fails, user will be prompted with the provided prompt.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.SlotValidation"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
