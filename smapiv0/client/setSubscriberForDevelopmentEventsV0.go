package client

import (
	subscriber_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscriber"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetSubscriberForDevelopmentEventsV0 Updates the properties of a subscriber.

	subscriberId - Unique identifier of the subscriber.
	updateSubscriberRequest - Defines the request body for updateSubscriber API.
*/
func (s *Client) SetSubscriberForDevelopmentEventsV0(subscriberId string, updateSubscriberRequest *subscriber_.UpdateSubscriberRequest) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v0/developmentEvents/subscribers/{subscriberId}")
	h.Path("subscriberId", subscriberId)
	h.Body = updateSubscriberRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Updates the properties of a subscriber.",
 "parameters": [
  {
   "description": "Unique identifier of the subscriber.",
   "format": "Amazon Common Identifier",
   "in": "path",
   "name": "subscriberId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the request body for updateSubscriber API.",
   "in": "body",
   "name": "UpdateSubscriberRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscriber.UpdateSubscriberRequest"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "Success."
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
 "x-operation-name": "setSubscriberForDevelopmentEventsV0"
}
*/
