package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	subscription_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscription"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateSubscriptionForDevelopmentEventsV0 Creates a new subscription for a subscriber. This needs to be authorized by the client/vendor who created the subscriber and the vendor who publishes the event.

	createSubscriptionRequest - Request body for createSubscription API.
*/
func (s *Client) CreateSubscriptionForDevelopmentEventsV0(createSubscriptionRequest *subscription_.CreateSubscriptionRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v0/developmentEvents/subscriptions")
	h.Body = createSubscriptionRequest
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
 "description": "Creates a new subscription for a subscriber. This needs to be authorized by the client/vendor who created the subscriber and the vendor who publishes the event.",
 "parameters": [
  {
   "description": "Request body for createSubscription API.",
   "in": "body",
   "name": "CreateSubscriptionRequest",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.CreateSubscriptionRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Created; Returns a URL to retrieve the subscription in 'Location' header.",
   "headers": {
    "Location": {
     "description": "Relative URL to retrieve the subscription.",
     "type": "string"
    }
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
 "x-operation-name": "createSubscriptionForDevelopmentEventsV0"
}
*/
