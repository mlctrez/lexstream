package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetPrivateDistributionAccountIdV1 Add an id to the private distribution accounts.

	skillId - The skill ID.
	stage - Stage for skill.
	id - ARN that a skill can be privately distributed to.
*/
func (s *Client) SetPrivateDistributionAccountIdV1(skillId string, stage string, id string) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/privateDistributionAccounts/{id}")
	h.Path("skillId", skillId)
	h.Path("stage", stage)
	h.Path("id", id)
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
 "description": "Add an id to the private distribution accounts.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage for skill.",
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "ARN that a skill can be privately distributed to.",
   "in": "path",
   "name": "id",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "Success.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only \"application/json\" supported.",
     "type": "string"
    }
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
 "x-operation-name": "setPrivateDistributionAccountIdV1"
}
*/
