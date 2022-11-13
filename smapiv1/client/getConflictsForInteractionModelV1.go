package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	conflictDetection_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/conflictDetection"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetConflictsForInteractionModelV1 This is a paginated API that retrieves results of conflict detection job for a specified interaction model.

	skillId - The skill ID.
	locale - The locale for the model requested e.g. en-GB, en-US, de-DE.
	stage - Stage of the interaction model.
	version - Version of interaction model. Use "~current" to get the model of the current version.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 100. If more results are present, the response will contain a nextToken and a _link.next href.
*/
func (s *Client) GetConflictsForInteractionModelV1(skillId string, locale string, stage string, version string, nextToken string, maxResults int) (response *conflictDetection_.GetConflictsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/interactionModel/locales/{locale}/versions/{version}/conflicts")
	h.Path("skillId", skillId)
	h.Path("locale", locale)
	h.Path("stage", stage)
	h.Path("version", version)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	response = &conflictDetection_.GetConflictsResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
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
 "description": "This is a paginated API that retrieves results of conflict detection job for a specified interaction model.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The locale for the model requested e.g. en-GB, en-US, de-DE.",
   "format": "languager-region; same as BCP-47 language tag format",
   "in": "path",
   "name": "locale",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage of the interaction model.",
   "enum": [
    "development"
   ],
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "Version of interaction model. Use \"~current\" to get the model of the current version.",
   "in": "path",
   "name": "version",
   "pattern": "^~current|[0-9]+$",
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
   "default": 100,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 100. If more results are present, the response will contain a nextToken and a _link.next href.",
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
   "description": "Get conflict detection results sucessfully.",
   "headers": {
    "Content-Type": {
     "description": "returned content type, only application/json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.GetConflictsResponse"
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
   "description": "There is no catalog defined for the catalogId.",
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
 "summary": "Retrieve conflict detection results for a specified interaction model.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getConflictsForInteractionModelV1"
}
*/
