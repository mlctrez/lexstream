package client

import (
	subscription_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscription"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetSubscriptionForDevelopmentEventsV0 Updates the mutable properties of a subscription. This needs to be authorized by the client/vendor who created the subscriber and the vendor who publishes the event. The subscriberId cannot be updated.

	subscriptionId - Unique identifier of the subscription.
	updateSubscriptionRequest - Request body for updateSubscription API.
*/
func (s *Client) SetSubscriptionForDevelopmentEventsV0(subscriptionId string, updateSubscriptionRequest *subscription_.UpdateSubscriptionRequest) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v0/developmentEvents/subscriptions/{subscriptionId}")
	h.Path("subscriptionId", subscriptionId)
	h.Body = updateSubscriptionRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Updates the mutable properties of a subscription. This needs to be authorized by the client/vendor who created the subscriber and the vendor who publishes the event. The subscriberId cannot be updated.",
 "parameters": [
  {
   "description": "Unique identifier of the subscription.",
   "format": "Amazon Common Identifier",
   "in": "path",
   "name": "subscriptionId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Request body for updateSubscription API.",
   "in": "body",
   "name": "UpdateSubscriptionRequest",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.UpdateSubscriptionRequest"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "No content."
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v0.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Defines the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v0.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "setSubscriptionForDevelopmentEventsV0"
}
*/
