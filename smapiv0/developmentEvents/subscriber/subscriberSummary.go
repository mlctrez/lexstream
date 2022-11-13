package subscriber

type SubscriberSummary struct {
	// Unique identifier of the subscriber resource.
	SubscriberId string `json:"subscriberId,omitempty"`
	// Name of the subscriber.
	Name   string            `json:"name,omitempty"`
	Status *SubscriberStatus `json:"status,omitempty"`
	// Client Id of the subscriber resource.
	ClientId string    `json:"clientId,omitempty"`
	Endpoint *Endpoint `json:"endpoint,omitempty"`
}

/*
{
 "properties": {
  "clientId": {
   "description": "Client Id of the subscriber resource.",
   "type": "string"
  },
  "endpoint": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.Endpoint"
  },
  "name": {
   "description": "Name of the subscriber.",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.SubscriberStatus",
   "x-isEnum": true
  },
  "subscriberId": {
   "description": "Unique identifier of the subscriber resource.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
