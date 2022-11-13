package evaluations

type Intent struct {
	Name               string                  `json,omitempty:"name"`
	ConfirmationStatus *ConfirmationStatusType `json,omitempty:"confirmationStatus"`
	/*
	   A map of key-value pairs that further describes what the user meant based on a predefined intent schema. The map can be empty.
	*/
	Slots map[string]Slot `json,omitempty:"slots"`
}

/*
{
 "properties": {
  "confirmationStatus": {
   "$ref": "#/definitions/v1.skill.evaluations.ConfirmationStatusType",
   "x-isEnum": true
  },
  "name": {
   "type": "string"
  },
  "slots": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.evaluations.Slot"
   },
   "description": "A map of key-value pairs that further describes what the user meant based on a predefined intent schema. The map can be empty.\n",
   "type": "object"
  }
 },
 "type": "object"
}
*/
