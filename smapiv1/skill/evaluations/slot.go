package evaluations

type Slot struct {
	Name               string                  `json,omitempty:"name"`
	Value              string                  `json,omitempty:"value"`
	ConfirmationStatus *ConfirmationStatusType `json,omitempty:"confirmationStatus"`
	Resolutions        *SlotResolutions        `json,omitempty:"resolutions"`
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
