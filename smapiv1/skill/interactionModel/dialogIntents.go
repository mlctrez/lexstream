package interactionmodel

type DialogIntents struct {
	// Name of the intent that has a dialog specification.
	Name string `json,omitempty:"name"`
	// Defines an intent-specific delegation strategy for this dialog intent. Overrides dialog-level setting.
	DelegationStrategy *DelegationStrategyType `json,omitempty:"delegationStrategy"`
	// List of slots that have dialog rules.
	Slots []*DialogSlotItems `json,omitempty:"slots"`
	// Describes whether confirmation of the intent is required.
	ConfirmationRequired bool                  `json,omitempty:"confirmationRequired"`
	Prompts              *DialogIntentsPrompts `json,omitempty:"prompts"`
}

/*
{
 "properties": {
  "confirmationRequired": {
   "description": "Describes whether confirmation of the intent is required.",
   "type": "boolean"
  },
  "delegationStrategy": {
   "$ref": "#/definitions/v1.skill.interactionModel.DelegationStrategyType",
   "description": "Defines an intent-specific delegation strategy for this dialog intent. Overrides dialog-level setting.",
   "x-isEnum": true
  },
  "name": {
   "description": "Name of the intent that has a dialog specification.",
   "type": "string"
  },
  "prompts": {
   "$ref": "#/definitions/v1.skill.interactionModel.DialogIntentsPrompts"
  },
  "slots": {
   "description": "List of slots that have dialog rules.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.DialogSlotItems"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
