package manifest

type EventName struct {
	EventName *EventNameType `json:"eventName,omitempty"`
}

/*
{
 "properties": {
  "eventName": {
   "$ref": "#/definitions/v1.skill.Manifest.EventNameType",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
