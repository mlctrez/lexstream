package interactionmodel

type DialogIntents struct {
	// Name of the intent that has a dialog specification.
	Name string `json:"name,omitempty"`
	// Defines an intent-specific delegation strategy for this dialog intent. Overrides dialog-level setting.
	DelegationStrategy *DelegationStrategyType `json:"delegationStrategy,omitempty"`
	// List of slots that have dialog rules.
	Slots []*DialogSlotItems `json:"slots,omitempty"`
	// Describes whether confirmation of the intent is required.
	ConfirmationRequired bool                  `json:"confirmationRequired,omitempty"`
	Prompts              *DialogIntentsPrompts `json:"prompts,omitempty"`
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
