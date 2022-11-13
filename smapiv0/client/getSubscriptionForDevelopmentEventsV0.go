package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	subscription_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscription"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSubscriptionForDevelopmentEventsV0 Returns information about a particular subscription. Both, the vendor who created the subscriber and the vendor who publishes the event can retrieve this resource with appropriate authorization.

	subscriptionId - Unique identifier of the subscription.
*/
func (s *Client) GetSubscriptionForDevelopmentEventsV0(subscriptionId string) (response *subscription_.SubscriptionInfo, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v0/developmentEvents/subscriptions/{subscriptionId}")
	h.Path("subscriptionId", subscriptionId)
	response = &subscription_.SubscriptionInfo{}
	h.Response = response
	h.ResponseType(400, &smapiv0.BadRequestError{})
	h.ResponseType(401, &smapiv0.Error{})
	h.ResponseType(403, &smapiv0.BadRequestError{})
	h.ResponseType(404, &smapiv0.Error{})
	h.ResponseType(429, &smapiv0.Error{})
	h.ResponseType(500, &smapiv0.Error{})
	h.ResponseType(503, &smapiv0.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Returns information about a particular subscription. Both, the vendor who created the subscriber and the vendor who publishes the event can retrieve this resource with appropriate authorization.",
 "parameters": [
  {
   "description": "Unique identifier of the subscription.",
   "format": "Amazon Common Identifier",
   "in": "path",
   "name": "subscriptionId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successful operation.",
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.SubscriptionInfo"
   }
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
 "x-operation-name": "getSubscriptionForDevelopmentEventsV0"
}
*/
