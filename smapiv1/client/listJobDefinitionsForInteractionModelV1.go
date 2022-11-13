package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	jobs_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/jobs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListJobDefinitionsForInteractionModelV1 Retrieve a list of jobs associated with the vendor.

	vendorId - The vendor ID.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
*/
func (s *Client) ListJobDefinitionsForInteractionModelV1(vendorId string, maxResults int, nextToken string) (response *jobs_.ListJobDefinitionsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs")
	h.Param("vendorId", vendorId)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	response = &jobs_.ListJobDefinitionsResponse{}
	h.Response = response
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Retrieve a list of jobs associated with the vendor.",
 "parameters": [
  {
   "description": "The vendor ID.",
   "in": "query",
   "name": "vendorId",
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
 "responses": {
  "200": {
   "description": "List of all jobs associated with the vendor.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.ListJobDefinitionsResponse"
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listJobDefinitionsForInteractionModelV1"
}
*/
