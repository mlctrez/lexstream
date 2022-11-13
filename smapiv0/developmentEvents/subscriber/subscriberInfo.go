package subscriber

/*
SubscriberInfo Information about the subscriber.
*/
type SubscriberInfo struct {
	// Unique identifier of the subscriber resource.
	SubscriberId string `json:"subscriberId"`
	// Name of the subscriber.
	Name     string    `json:"name"`
	Endpoint *Endpoint `json:"endpoint"`
}

/*
{
 "description": "Information about the subscriber.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.Endpoint"
  },
  "name": {
   "description": "Name of the subscriber.",
   "type": "string"
  },
  "subscriberId": {
   "description": "Unique identifier of the subscriber resource.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
