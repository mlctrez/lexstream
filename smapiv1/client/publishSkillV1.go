package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	publication_ "github.com/mlctrez/lexstream/smapiv1/skill/publication"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
PublishSkillV1 If the skill is in certified stage, initiate publishing immediately or set a date at which the skill can publish at.

	skillId - The skill ID.
	accept_Language - User's locale/language in context.
	publishSkillRequest - Defines the request body for publish skill API.
*/
func (s *Client) PublishSkillV1(skillId string, accept_Language string, publishSkillRequest *publication_.PublishSkillRequest) (response *publication_.SkillPublicationResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/publications")
	h.Path("skillId", skillId)
	h.Header("Accept-Language", accept_Language)
	h.Body = publishSkillRequest
	response = &publication_.SkillPublicationResponse{}
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
 "description": "If the skill is in certified stage, initiate publishing immediately or set a date at which the skill can publish at.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "User's locale/language in context.",
   "in": "header",
   "name": "Accept-Language",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the request body for publish skill API.",
   "in": "body",
   "name": "publishSkillRequest",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v1.skill.publication.PublishSkillRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Successfully processed skill publication request.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.publication.SkillPublicationResponse"
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
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
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
 "x-operation-name": "publishSkillV1"
}
*/
