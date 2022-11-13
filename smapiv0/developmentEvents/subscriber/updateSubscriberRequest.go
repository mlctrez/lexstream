package subscriber

type UpdateSubscriberRequest struct {
	// Name of the subscriber.
	Name     string    `json,omitempty:"name"`
	Endpoint *Endpoint `json,omitempty:"endpoint"`
}

/*
{
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.Endpoint"
  },
  "name": {
   "description": "Name of the subscriber.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
