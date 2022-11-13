package evaluations

type SlotsProps struct {
	Name               string              `json:"name"`
	Value              string              `json:"value"`
	ConfirmationStatus *ConfirmationStatus `json:"confirmationStatus"`
	Resolutions        *Resolutions        `json:"resolutions"`
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
