package subscriber

type UpdateSubscriberRequest struct {
	// Name of the subscriber.
	Name     string    `json:"name"`
	Endpoint *Endpoint `json:"endpoint"`
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
