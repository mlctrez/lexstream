package evaluations

type SlotsProps struct {
	Name               string              `json,omitempty:"name"`
	Value              string              `json,omitempty:"value"`
	ConfirmationStatus *ConfirmationStatus `json,omitempty:"confirmationStatus"`
	Resolutions        *Resolutions        `json,omitempty:"resolutions"`
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
  "resolutions": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Resolutions"
  },
  "value": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
