package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	subscription_ "github.com/mlctrez/lexstream/smapiv0/developmentEvents/subscription"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListSubscriptionsForDevelopmentEventsV0 Lists all the subscriptions for a vendor/subscriber depending on the query parameter.

	vendorId - The vendor ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	subscriberId - Unique identifier of the subscriber. If this query parameter is provided, the list would be filtered by the owning subscriberId.
*/
func (s *Client) ListSubscriptionsForDevelopmentEventsV0(vendorId string, nextToken string, maxResults int, subscriberId string) (response *subscription_.ListSubscriptionsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v0/developmentEvents/subscriptions")
	h.Param("vendorId", vendorId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	h.Param("subscriberId", subscriberId)
	response = &subscription_.ListSubscriptionsResponse{}
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
 "description": "Lists all the subscriptions for a vendor/subscriber depending on the query parameter.",
 "parameters": [
  {
   "description": "The vendor ID.",
   "in": "query",
   "name": "vendorId",
   "required": true,
   "type": "string"
  },
  {
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.",
   "in": "query",
   "maximum": 50,
   "minimum": 1,
   "multipleOf": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "description": "Unique identifier of the subscriber. If this query parameter is provided, the list would be filtered by the owning subscriberId.",
   "format": "Amazon Common Identifier",
   "in": "query",
   "name": "subscriberId",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successful operation.",
   "schema": {
    "$ref": "#/definitions/v0.developmentEvents.subscription.ListSubscriptionsResponse"
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
 "x-operation-name": "listSubscriptionsForDevelopmentEventsV0"
}
*/
