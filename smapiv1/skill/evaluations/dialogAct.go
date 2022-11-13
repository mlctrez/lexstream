package evaluations

/*
DialogAct A representation of question prompts to the user for multi-turn, which requires user to fill a slot value, or confirm a slot value, or confirm an intent.
*/
type DialogAct struct {
	Type_ *DialogActType `json,omitempty:"type"`
	// The name of the target slot that needs to be filled or confirmed for a dialogAct
	TargetSlot string `json,omitempty:"targetSlot"`
}

/*
{
 "description": "A representation of question prompts to the user for multi-turn, which requires user to fill a slot value, or confirm a slot value, or confirm an intent.\n",
 "properties": {
  "targetSlot": {
   "description": "The name of the target slot that needs to be filled or confirmed for a dialogAct",
   "type": "string"
  },
  "type": {
   "$ref": "#/definitions/v1.skill.evaluations.DialogActType",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
