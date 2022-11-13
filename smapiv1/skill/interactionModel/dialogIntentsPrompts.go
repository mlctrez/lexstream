package interactionmodel

/*
DialogIntentsPrompts Collection of prompts for this intent.
*/
type DialogIntentsPrompts struct {
	// Enum value in the dialog_slots map to reference the elicitation prompt id.
	Elicitation string `json:"elicitation,omitempty"`
	// Enum value in the dialog_slots map to reference the confirmation prompt id.
	Confirmation string `json:"confirmation,omitempty"`
}

/*
{
 "description": "Collection of prompts for this intent.",
 "properties": {
  "confirmation": {
   "description": "Enum value in the dialog_slots map to reference the confirmation prompt id.",
   "type": "string"
  },
  "elicitation": {
   "description": "Enum value in the dialog_slots map to reference the elicitation prompt id.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
