package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	jobs_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/jobs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListJobExecutionsForInteractionModelV1 List the execution history associated with the job definition, with default sortField to be the executions' timestamp.

	jobId - The identifier for dynamic jobs.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	sortDirection - Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.
*/
func (s *Client) ListJobExecutionsForInteractionModelV1(jobId string, maxResults int, nextToken string, sortDirection string) (response *jobs_.GetExecutionsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs/{jobId}/executions")
	h.Path("jobId", jobId)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	h.Param("sortDirection", sortDirection)
	response = &jobs_.GetExecutionsResponse{}
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
 "description": "List the execution history associated with the job definition, with default sortField to be the executions' timestamp.",
 "parameters": [
  {
   "description": "The identifier for dynamic jobs.",
   "in": "path",
   "name": "jobId",
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
  },
  {
   "description": "Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.",
   "in": "query",
   "name": "sortDirection",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Retrun list of executions associated with the job definition.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.GetExecutionsResponse"
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
 "x-operation-name": "listJobExecutionsForInteractionModelV1"
}
*/
