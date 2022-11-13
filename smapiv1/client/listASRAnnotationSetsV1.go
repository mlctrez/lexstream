package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListASRAnnotationSetsV1 API which requests all the ASR annotation sets for a skill. Returns the annotation set id and properties for each ASR annotation set. Supports paging of results.

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a paginationContext.
*/
func (s *Client) ListASRAnnotationSetsV1(skillId string, nextToken string, maxResults int) (response *annotationSets_.ListASRAnnotationSetsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	response = &annotationSets_.ListASRAnnotationSetsResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which requests all the ASR annotation sets for a skill. Returns the annotation set id and properties for each ASR annotation set. Supports paging of results.\n",
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
   "default": 1000,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a paginationContext.\n",
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
   "description": "ASR annotation sets metadata are returned.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/hal+json or application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.ListASRAnnotationSetsResponse"
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
 "summary": "List ASR annotation sets metadata for a given skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listASRAnnotationSetsV1"
}
*/
