package evaluations

type Slot struct {
	Name               string                  `json:"name"`
	Value              string                  `json:"value"`
	ConfirmationStatus *ConfirmationStatusType `json:"confirmationStatus"`
	Resolutions        *SlotResolutions        `json:"resolutions"`
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
  "resolutions": {
   "$ref": "#/definitions/v1.skill.evaluations.SlotResolutions"
  },
  "value": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
