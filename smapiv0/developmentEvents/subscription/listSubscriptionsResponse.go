package subscription

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

type ListSubscriptionsResponse struct {
	Links     *smapiv0.Links `json,omitempty:"_links"`
	NextToken string         `json,omitempty:"nextToken"`
	// List of subscription summaries.
	Subscriptions []*SubscriptionSummary `json,omitempty:"subscriptions"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v0.Links"
  },
  "nextToken": {
   "type": "string"
  },
  "subscriptions": {
   "description": "List of subscription summaries.",
   "items": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.SubscriptionSummary"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
