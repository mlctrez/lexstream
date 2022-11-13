package subscription

type SubscriptionInfo struct {
	// Name of the subscription.
	Name string `json,omitempty:"name"`
	// Unique identifier of the subscription resource.
	SubscriptionId string `json,omitempty:"subscriptionId"`
	// Subscriber Id of the event-receiver.
	SubscriberId string `json,omitempty:"subscriberId"`
	// Vendor Id of the event-publisher.
	VendorId string `json,omitempty:"vendorId"`
	// The list of events that the subscriber should be notified for.
	Events []*Event `json,omitempty:"events"`
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
   "description": "Subscriber Id of the event-receiver.",
   "type": "string"
  },
  "subscriptionId": {
   "description": "Unique identifier of the subscription resource.",
   "type": "string"
  },
  "vendorId": {
   "description": "Vendor Id of the event-publisher.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
