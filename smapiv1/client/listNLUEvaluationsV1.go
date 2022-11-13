package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListNLUEvaluationsV1 API which requests recently run nlu evaluations started by a vendor for a skill. Returns the evaluation id and some of the parameters used to start the evaluation. Developers can filter the results using locale and stage. Supports paging of results.

	skillId - The skill ID.
	locale - filter to evaluations started using this locale
	stage - filter to evaluations started using this stage
	annotationId - filter to evaluations started using this annotationId
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 10. If more results are present, the response will contain a nextToken and a _link.next href.
*/
func (s *Client) ListNLUEvaluationsV1(skillId string, locale string, stage string, annotationId string, nextToken string, maxResults int) (response *evaluations_.ListNLUEvaluationsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/nluEvaluations")
	h.Path("skillId", skillId)
	h.Param("locale", locale)
	h.Param("stage", stage)
	h.Param("annotationId", annotationId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	response = &evaluations_.ListNLUEvaluationsResponse{}
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
 "description": "API which requests recently run nlu evaluations started by a vendor for a skill. Returns the evaluation id and some of the parameters used to start the evaluation. Developers can filter the results using locale and stage. Supports paging of results.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "filter to evaluations started using this locale",
   "in": "query",
   "name": "locale",
   "required": false,
   "type": "string"
  },
  {
   "description": "filter to evaluations started using this stage",
   "in": "query",
   "name": "stage",
   "required": false,
   "type": "string"
  },
  {
   "description": "filter to evaluations started using this annotationId",
   "in": "query",
   "name": "annotationId",
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
   "default": 10,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 10. If more results are present, the response will contain a nextToken and a _link.next href.\n",
   "in": "query",
   "maximum": 100,
   "minimum": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  }
 ],
 "responses": {
  "200": {
   "description": "Evaluations are returned.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/hal+json or application/json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.ListNLUEvaluationsResponse"
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
 "summary": "List nlu evaluations run for a skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listNLUEvaluationsV1"
}
*/
