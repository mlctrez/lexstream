package evaluations

/*
MultiTurn Included when the selected intent has dialog defined and the dialog is not completed.  To continue the dialog, provide the value of the token in the multiTurnToken field in the next request.
*/
type MultiTurn struct {
	DialogAct *DialogAct `json,omitempty:"dialogAct"`
	// Opaque string which contains multi-turn related context.
	Token string `json,omitempty:"token"`
	// A sample prompt defined in the dialog model for each DialogAct.
	Prompt string `json,omitempty:"prompt"`
}

/*
{
 "description": "Included when the selected intent has dialog defined and the dialog is not completed.  To continue the dialog, provide the value of the token in the multiTurnToken field in the next request.\n",
 "properties": {
  "dialogAct": {
   "$ref": "#/definitions/v1.skill.evaluations.DialogAct"
  },
  "prompt": {
   "description": "A sample prompt defined in the dialog model for each DialogAct.",
   "type": "string"
  },
  "token": {
   "description": "Opaque string which contains multi-turn related context.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
