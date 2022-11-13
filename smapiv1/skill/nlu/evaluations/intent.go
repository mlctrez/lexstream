package evaluations

type Intent struct {
	Name               string                `json:"name,omitempty"`
	ConfirmationStatus *ConfirmationStatus   `json:"confirmationStatus,omitempty"`
	Slots              map[string]SlotsProps `json:"slots,omitempty"`
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
