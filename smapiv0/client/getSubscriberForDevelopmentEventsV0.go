package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	subscriber_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscriber"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSubscriberForDevelopmentEventsV0 Returns information about specified subscriber.

	subscriberId - Unique identifier of the subscriber.
*/
func (s *Client) GetSubscriberForDevelopmentEventsV0(subscriberId string) (response *subscriber_.SubscriberInfo, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v0/developmentEvents/subscribers/{subscriberId}")
	h.Path("subscriberId", subscriberId)
	response = &subscriber_.SubscriberInfo{}
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
 "description": "Returns information about specified subscriber.",
 "parameters": [
  {
   "description": "Unique identifier of the subscriber.",
   "format": "Amazon Common Identifier",
   "in": "path",
   "name": "subscriberId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successful operation.",
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscriber.SubscriberInfo"
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
 "x-operation-name": "getSubscriberForDevelopmentEventsV0"
}
*/
