package client

import (
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListASREvaluationsResultsV1 Paginated API which returns the test case results of an evaluation. This should be considered the "expensive" operation while GetAsrEvaluationsStatus is "cheap".

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	evaluationId - Identifier of the evaluation.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a nextToken.

	status - query parameter used to filter evaluation result status.
	 * `PASSED` - filter evaluation result status of `PASSED`
	 * `FAILED` - filter evaluation result status of `FAILED`
*/
func (s *Client) ListASREvaluationsResultsV1(skillId string, nextToken string, evaluationId string, maxResults int, status string) (response *evaluations_.GetAsrEvaluationsResultsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/asrEvaluations/{evaluationId}/results")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Path("evaluationId", evaluationId)
	h.Param("maxResults", maxResults)
	h.Param("status", status)
	response = &evaluations_.GetAsrEvaluationsResultsResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Paginated API which returns the test case results of an evaluation. This should be considered the \"expensive\" operation while GetAsrEvaluationsStatus is \"cheap\".\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
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
   "description": "Identifier of the evaluation.",
   "in": "path",
   "name": "evaluationId",
   "required": true,
   "type": "string"
  },
  {
   "default": 1000,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a nextToken.\n",
   "in": "query",
   "maximum": 1000,
   "minimum": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "description": "query parameter used to filter evaluation result status.\n  * `PASSED` - filter evaluation result status of `PASSED`\n  * `FAILED` - filter evaluation result status of `FAILED`\n",
   "enum": [
    "PASSED",
    "FAILED"
   ],
   "in": "query",
   "name": "status",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Evaluation exists and its status is queryable.",
   "headers": {
    "Content-Type": {
     "description": "returned content type, - application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.GetAsrEvaluationsResultsResponse"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Define the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
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
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "default": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "List results for a completed Evaluation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listASREvaluationsResultsV1"
}
*/
