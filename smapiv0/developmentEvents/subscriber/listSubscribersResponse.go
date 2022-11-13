package subscriber

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

type ListSubscribersResponse struct {
	Links     *smapiv0.Links `json,omitempty:"_links"`
	NextToken string         `json,omitempty:"nextToken"`
	// List containing subscriber summary.
	Subscribers []*SubscriberSummary `json,omitempty:"subscribers"`
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
  "subscribers": {
   "description": "List containing subscriber summary.",
   "items": {
    "$ref": "#/definitions/v0.developmentEvents.subscriber.SubscriberSummary"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
