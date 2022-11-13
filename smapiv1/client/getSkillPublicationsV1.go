package client

import (
	publication_ "github.com/mlctrez/lexstream/smapiv1/skill/publication"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSkillPublicationsV1 Retrieves the latest skill publishing details of the certified stage of the skill. The publishesAtDate and
status of skill publishing.

	skillId - The skill ID.
	accept_Language - User's locale/language in context.
*/
func (s *Client) GetSkillPublicationsV1(skillId string, accept_Language string) (response *publication_.SkillPublicationResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/publications/~latest")
	h.Path("skillId", skillId)
	h.Header("Accept-Language", accept_Language)
	response = &publication_.SkillPublicationResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Retrieves the latest skill publishing details of the certified stage of the skill. The publishesAtDate and\nstatus of skill publishing.\n",
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
  }
 ],
 "responses": {
  "200": {
   "description": "Successfully retrieved latest skill publication information.",
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
 "x-operation-name": "getSkillPublicationsV1"
}
*/
