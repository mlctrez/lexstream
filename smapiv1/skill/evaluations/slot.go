package evaluations

type Slot struct {
	Name               string                  `json:"name,omitempty"`
	Value              string                  `json:"value,omitempty"`
	ConfirmationStatus *ConfirmationStatusType `json:"confirmationStatus,omitempty"`
	Resolutions        *SlotResolutions        `json:"resolutions,omitempty"`
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
