package manifest

type EventName struct {
	EventName *EventNameType `json,omitempty:"eventName"`
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
