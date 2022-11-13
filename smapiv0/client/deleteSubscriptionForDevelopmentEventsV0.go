package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
DeleteSubscriptionForDevelopmentEventsV0 Deletes a particular subscription. Both, the vendor who created the subscriber and the vendor who publishes the event can delete this resource with appropriate authorization.

	subscriptionId - Unique identifier of the subscription.
*/
func (s *Client) DeleteSubscriptionForDevelopmentEventsV0(subscriptionId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v0/developmentEvents/subscriptions/{subscriptionId}")
	h.Path("subscriptionId", subscriptionId)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Deletes a particular subscription. Both, the vendor who created the subscriber and the vendor who publishes the event can delete this resource with appropriate authorization.",
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
  "204": {
   "description": "Successful operation."
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
 "x-operation-name": "deleteSubscriptionForDevelopmentEventsV0"
}
*/
