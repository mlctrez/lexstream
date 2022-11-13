package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	smartHomeEvaluation_ "github.com/mlctrez/lexstream/smapiv1/smartHomeEvaluation"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListSmarthomeCapabilityTestPlansV1 List all the test plan names and ids for a given skill ID.

	skillId - The skill ID.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
*/
func (s *Client) ListSmarthomeCapabilityTestPlansV1(skillId string, maxResults int, nextToken string) (response *smartHomeEvaluation_.ListSHCapabilityTestPlansResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/smartHome/testing/capabilityTestPlans")
	h.Path("skillId", skillId)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	response = &smartHomeEvaluation_.ListSHCapabilityTestPlansResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.BadRequestError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.BadRequestError{})
	h.ResponseType(429, &smapiv1.BadRequestError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "List all the test plan names and ids for a given skill ID.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
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
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  }
 ],
 "produces": [
  "application/json"
 ],
 "responses": {
  "200": {
   "description": "Successfully got the list of test plans.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.ListSHCapabilityTestPlansResponse"
   }
  },
  "400": {
   "description": "Bad Request. Returned when the request payload is malformed or when, at least, one required property is missing or invalid.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "403": {
   "description": "API user does not have permission or is currently in a state that does not allow calls to this API.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill, test plan, or evaluation does not exist. The error response will contain a description that indicates the specific resource type that was not found.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "429": {
   "description": "Exceeded the permitted request limit. Throttling criteria includes total requests, per API and CustomerId.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "default": {
   "description": "Internal server error.\n",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "List all the test plan names and ids for a given skill ID.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listSmarthomeCapabilityTestPlansV1"
}
*/
