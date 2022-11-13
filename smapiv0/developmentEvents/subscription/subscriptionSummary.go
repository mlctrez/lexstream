package subscription

type SubscriptionSummary struct {
	// Name of the subscription.
	Name string `json:"name,omitempty"`
	// Unique identifier of the subscription resource.
	SubscriptionId string `json:"subscriptionId,omitempty"`
	// Subscriber Id of the event-reciever.
	SubscriberId string `json:"subscriberId,omitempty"`
	// VendorId of the event-publisher.
	VendorId string `json:"vendorId,omitempty"`
	// The list of events that the subscriber should be notified for.
	Events []*Event `json:"events,omitempty"`
}

/*
{
 "properties": {
  "events": {
   "description": "The list of events that the subscriber should be notified for.",
   "items": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.Event"
   },
   "type": "array"
  },
  "name": {
   "description": "Name of the subscription.",
   "type": "string"
  },
  "subscriberId": {
   "description": "Subscriber Id of the event-reciever.",
   "type": "string"
  },
  "subscriptionId": {
   "description": "Unique identifier of the subscription resource.",
   "type": "string"
  },
  "vendorId": {
   "description": "VendorId of the event-publisher.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
