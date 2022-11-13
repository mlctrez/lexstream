package interactionmodel

/*
DialogPrompts Dialog prompts associated with this slot i.e. for elicitation and/or confirmation.
*/
type DialogPrompts struct {
	// Reference to a prompt-id to use If this slot value is missing.
	Elicitation string `json,omitempty:"elicitation"`
	// Reference to a prompt-id to use to confirm the slots value.
	Confirmation string `json,omitempty:"confirmation"`
}

/*
{
 "description": "Dialog prompts associated with this slot i.e. for elicitation and/or confirmation.",
 "properties": {
  "confirmation": {
   "description": "Reference to a prompt-id to use to confirm the slots value.",
   "type": "string"
  },
  "elicitation": {
   "description": "Reference to a prompt-id to use If this slot value is missing.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
