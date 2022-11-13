package subscription

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

type ListSubscriptionsResponse struct {
	Links     *smapiv0.Links `json:"_links,omitempty"`
	NextToken string         `json:"nextToken,omitempty"`
	// List of subscription summaries.
	Subscriptions []*SubscriptionSummary `json:"subscriptions,omitempty"`
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
