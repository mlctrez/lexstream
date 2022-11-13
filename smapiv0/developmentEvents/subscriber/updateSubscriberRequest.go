package subscriber

type UpdateSubscriberRequest struct {
	// Name of the subscriber.
	Name     string    `json:"name,omitempty"`
	Endpoint *Endpoint `json:"endpoint,omitempty"`
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
