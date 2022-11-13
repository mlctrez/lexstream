package evaluations

type Intent struct {
	Name               string                `json,omitempty:"name"`
	ConfirmationStatus *ConfirmationStatus   `json,omitempty:"confirmationStatus"`
	Slots              map[string]SlotsProps `json,omitempty:"slots"`
}

/*
{
 "properties": {
  "confirmationStatus": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.ConfirmationStatus",
   "x-isEnum": true
  },
  "name": {
   "type": "string"
  },
  "slots": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.SlotsProps"
   },
   "type": "object"
  }
 },
 "type": "object"
}
*/
