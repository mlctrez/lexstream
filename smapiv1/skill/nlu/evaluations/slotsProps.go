package evaluations

type SlotsProps struct {
	Name               string              `json:"name,omitempty"`
	Value              string              `json:"value,omitempty"`
	ConfirmationStatus *ConfirmationStatus `json:"confirmationStatus,omitempty"`
	Resolutions        *Resolutions        `json:"resolutions,omitempty"`
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
