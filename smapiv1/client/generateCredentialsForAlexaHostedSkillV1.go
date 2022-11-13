package client

import (
	AlexaHosted_ "github.com/mlctrez/lexstream/smapiv1/skill/AlexaHosted"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GenerateCredentialsForAlexaHostedSkillV1 Generates hosted skill repository credentials to access the hosted skill repository.

	skillId - The skill ID.
	hostedSkillRepositoryCredentialsRequest - defines the request body for hosted skill repository credentials
*/
func (s *Client) GenerateCredentialsForAlexaHostedSkillV1(skillId string, hostedSkillRepositoryCredentialsRequest *AlexaHosted_.HostedSkillRepositoryCredentialsRequest) (response *AlexaHosted_.HostedSkillRepositoryCredentialsList, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/alexaHosted/repository/credentials/generate")
	h.Path("skillId", skillId)
	h.Body = hostedSkillRepositoryCredentialsRequest
	response = &AlexaHosted_.HostedSkillRepositoryCredentialsList{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Generates hosted skill repository credentials to access the hosted skill repository.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "defines the request body for hosted skill repository credentials",
   "in": "body",
   "name": "hostedSkillRepositoryCredentialsRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepositoryCredentialsRequest"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Response contains the hosted skill repository credentials",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only \"application/json\" supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepositoryCredentialsList"
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
 "x-operation-name": "generateCredentialsForAlexaHostedSkillV1"
}
*/
