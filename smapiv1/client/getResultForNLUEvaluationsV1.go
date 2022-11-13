package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetResultForNLUEvaluationsV1 Paginated API which returns the test case results of an evaluation. This should be considered the 'expensive' operation while getNluEvaluation is 'cheap'.

	skillId - The skill ID.
	evaluationId - Identifier of the evaluation.
	sortfield -
	testCaseStatus - only returns test cases with this status
	actualIntentName - only returns test cases with intents which resolve to this intent
	expectedIntentName - only returns test cases with intents which are expected to be this intent
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a nextToken and a _link.next href.
*/
func (s *Client) GetResultForNLUEvaluationsV1(skillId string, evaluationId string, sortfield string, testCaseStatus string, actualIntentName string, expectedIntentName string, nextToken string, maxResults int) (response *evaluations_.GetNLUEvaluationResultsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/nluEvaluations/{evaluationId}/results")
	h.Path("skillId", skillId)
	h.Path("evaluationId", evaluationId)
	h.Param("sortfield", sortfield)
	h.Param("testCaseStatus", testCaseStatus)
	h.Param("actualIntentName", actualIntentName)
	h.Param("expectedIntentName", expectedIntentName)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	response = &evaluations_.GetNLUEvaluationResultsResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Paginated API which returns the test case results of an evaluation. This should be considered the 'expensive' operation while getNluEvaluation is 'cheap'.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
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
   "enum": [
    "STATUS",
    "ACTUAL_INTENT",
    "EXPECTED_INTENT"
   ],
   "in": "query",
   "name": "sort.field",
   "required": false,
   "type": "string"
  },
  {
   "description": "only returns test cases with this status",
   "enum": [
    "PASSED",
    "FAILED"
   ],
   "in": "query",
   "name": "testCaseStatus",
   "required": false,
   "type": "string"
  },
  {
   "description": "only returns test cases with intents which resolve to this intent",
   "in": "query",
   "name": "actualIntentName",
   "required": false,
   "type": "string"
  },
  {
   "description": "only returns test cases with intents which are expected to be this intent",
   "in": "query",
   "name": "expectedIntentName",
   "required": false,
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
   "default": 1000,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a nextToken and a _link.next href.\n",
   "in": "query",
   "maximum": 1000,
   "minimum": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  }
 ],
 "responses": {
  "200": {
   "description": "Evaluation exists and its status is queryable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, - application/json or application/hal+json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.GetNLUEvaluationResultsResponse"
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
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Get test case results for a completed Evaluation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getResultForNLUEvaluationsV1"
}
*/
