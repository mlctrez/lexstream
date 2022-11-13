package client

import (
	smartHomeEvaluation_ "github.com/mlctrez/lexstream/smapiv1/smartHomeEvaluation"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListSmarthomeCapabilityEvaluationsV1 List Smart Home capability evaluation runs for a skill.

	skillId - The skill ID.
	stage - The stage of the skill to be used for evaluation. An error will be returned if this skill stage is not enabled on the account used for evaluation.
	startTimestampFrom - The begnning of the start time to query evaluation result.
	startTimestampTo - The end of the start time to query evaluation result.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
*/
func (s *Client) ListSmarthomeCapabilityEvaluationsV1(skillId string, stage string, startTimestampFrom string, startTimestampTo string, maxResults int, nextToken string) (response *smartHomeEvaluation_.ListSHCapabilityEvaluationsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/smartHome/testing/capabilityEvaluations")
	h.Path("skillId", skillId)
	h.Param("stage", stage)
	h.Param("startTimestampFrom", startTimestampFrom)
	h.Param("startTimestampTo", startTimestampTo)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	response = &smartHomeEvaluation_.ListSHCapabilityEvaluationsResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "List Smart Home capability evaluation runs for a skill.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The stage of the skill to be used for evaluation. An error will be returned if this skill stage is not enabled on the account used for evaluation.",
   "enum": [
    "development",
    "live"
   ],
   "in": "query",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "The begnning of the start time to query evaluation result.",
   "format": "date-time",
   "in": "query",
   "name": "startTimestampFrom",
   "required": false,
   "type": "string"
  },
  {
   "description": "The end of the start time to query evaluation result.",
   "format": "date-time",
   "in": "query",
   "name": "startTimestampTo",
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
   "description": "Successfully retrieved the evaluation infomation.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.ListSHCapabilityEvaluationsResponse"
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
 "summary": "List Smart Home capability evaluation runs for a skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listSmarthomeCapabilityEvaluationsV1"
}
*/
