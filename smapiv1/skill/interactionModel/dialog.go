package interactionmodel

/*
Dialog Defines dialog rules e.g. slot elicitation and validation, intent chaining etc.
*/
type Dialog struct {
	// Defines a delegation strategy for the dialogs in this dialog model.
	DelegationStrategy *DelegationStrategyType `json,omitempty:"delegationStrategy"`
	// List of intents that have dialog rules associated with them. Dialogs can also span multiple intents.
	Intents []*DialogIntents `json,omitempty:"intents"`
}

/*
{
 "description": "Defines dialog rules e.g. slot elicitation and validation, intent chaining etc.",
 "properties": {
  "delegationStrategy": {
   "$ref": "#/definitions/v1.skill.interactionModel.DelegationStrategyType",
   "description": "Defines a delegation strategy for the dialogs in this dialog model.",
   "x-isEnum": true
  },
  "intents": {
   "description": "List of intents that have dialog rules associated with them. Dialogs can also span multiple intents.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.DialogIntents"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
