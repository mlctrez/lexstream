package subscription

type CreateSubscriptionRequest struct {
	// Name of the subscription.
	Name string `json,omitempty:"name"`
	// The list of events that the subscriber should be notified for.
	Events []*Event `json,omitempty:"events"`
	// The vendorId of the event publisher.
	VendorId string `json,omitempty:"vendorId"`
	// The id of the subscriber that would receive the events.
	SubscriberId string `json,omitempty:"subscriberId"`
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
   "description": "The id of the subscriber that would receive the events.",
   "type": "string"
  },
  "vendorId": {
   "description": "The vendorId of the event publisher.",
   "type": "string"
  }
 },
 "required": [
  "events",
  "name",
  "subscriberId",
  "vendorId"
 ],
 "type": "object"
}
*/
