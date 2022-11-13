package manifest

type EventPublications struct {
	// Name of the event to publish.
	EventName string `json:"eventName,omitempty"`
}

/*
{
 "properties": {
  "eventName": {
   "description": "Name of the event to publish.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
