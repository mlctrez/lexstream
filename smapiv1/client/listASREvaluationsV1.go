package client

import (
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListASREvaluationsV1 API that allows developers to get historical ASR evaluations they run before.

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	locale - locale in bcp 47 format. Used to filter results with the specified locale. If omitted, the response would include all evaluations regardless of what locale was used in the evaluation
	stage - Query parameter used to filter evaluations with specified skill stage.
	 * `development` - skill in `development` stage
	 * `live` - skill in `live` stage

	annotationSetId - filter to evaluations started using this annotationSetId
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a nextToken.
*/
func (s *Client) ListASREvaluationsV1(skillId string, nextToken string, locale string, stage string, annotationSetId string, maxResults int) (response *evaluations_.ListAsrEvaluationsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/asrEvaluations")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Param("locale", locale)
	h.Param("stage", stage)
	h.Param("annotationSetId", annotationSetId)
	h.Param("maxResults", maxResults)
	response = &evaluations_.ListAsrEvaluationsResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API that allows developers to get historical ASR evaluations they run before.\n",
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
   "description": "locale in bcp 47 format. Used to filter results with the specified locale. If omitted, the response would include all evaluations regardless of what locale was used in the evaluation",
   "format": "locale",
   "in": "query",
   "name": "locale",
   "required": false,
   "type": "string"
  },
  {
   "description": "Query parameter used to filter evaluations with specified skill stage.\n  * `development` - skill in `development` stage\n  * `live` - skill in `live` stage\n",
   "enum": [
    "development",
    "live"
   ],
   "in": "query",
   "name": "stage",
   "required": false,
   "type": "string"
  },
  {
   "description": "filter to evaluations started using this annotationSetId",
   "in": "query",
   "name": "annotationSetId",
   "required": false,
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
  }
 ],
 "responses": {
  "200": {
   "description": "Evaluations are returned.",
   "headers": {
    "Content-Type": {
     "description": "returned content type or application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.ListAsrEvaluationsResponse"
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
 "summary": "List asr evaluations run for a skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listASREvaluationsV1"
}
*/
