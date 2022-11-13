package subscription

type UpdateSubscriptionRequest struct {
	// Name of the subscription.
	Name string `json:"name"`
	// The list of events that the subscriber should be notified for.
	Events []*Event `json:"events"`
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
  }
 },
 "required": [
  "events",
  "name"
 ],
 "type": "object"
}
*/
