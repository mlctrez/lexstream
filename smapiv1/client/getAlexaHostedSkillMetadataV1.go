package client

import (
	AlexaHosted_ "github.com/mlctrez/lexstream/smapiv1/skill/AlexaHosted"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetAlexaHostedSkillMetadataV1 Get Alexa hosted skill's metadata

	skillId - The skill ID.
*/
func (s *Client) GetAlexaHostedSkillMetadataV1(skillId string) (response *AlexaHosted_.HostedSkillMetadata, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/alexaHosted")
	h.Path("skillId", skillId)
	response = &AlexaHosted_.HostedSkillMetadata{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get Alexa hosted skill's metadata",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "response contains the Alexa hosted skill's metadata",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only \"application/json\" supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillMetadata"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. Authorization Url is invalid",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "defines the authentication method that should be used to gain access to a resource.",
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
 "x-operation-name": "getAlexaHostedSkillMetadataV1"
}
*/
