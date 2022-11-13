package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	subscriber_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscriber"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateSubscriberForDevelopmentEventsV0 Creates a new subscriber resource for a vendor.

	createSubscriberRequest - Defines the request body for createSubscriber API.
*/
func (s *Client) CreateSubscriberForDevelopmentEventsV0(createSubscriberRequest *subscriber_.CreateSubscriberRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v0/developmentEvents/subscribers")
	h.Body = createSubscriberRequest
	h.ResponseType(400, &smapiv0.BadRequestError{})
	h.ResponseType(401, &smapiv0.Error{})
	h.ResponseType(429, &smapiv0.Error{})
	h.ResponseType(500, &smapiv0.Error{})
	h.ResponseType(503, &smapiv0.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new subscriber resource for a vendor.",
 "parameters": [
  {
   "description": "Defines the request body for createSubscriber API.",
   "in": "body",
   "name": "CreateSubscriberRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscriber.CreateSubscriberRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Created. Returns a URL to retrieve the subscriber in 'Location' header.",
   "headers": {
    "Location": {
     "description": "Relative URL to retrieve the subscriber.",
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
 "x-operation-name": "createSubscriberForDevelopmentEventsV0"
}
*/
