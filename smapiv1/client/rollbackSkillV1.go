package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
RollbackSkillV1 Submit a target skill version to rollback to. Only one rollback or publish operation can be outstanding for a given skillId.

	skillId - The skill ID.
	createRollbackRequest - defines the request body to create a rollback request
*/
func (s *Client) RollbackSkillV1(skillId string, createRollbackRequest *skill_.CreateRollbackRequest) (response *skill_.CreateRollbackResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/rollbacks")
	h.Path("skillId", skillId)
	h.Body = createRollbackRequest
	response = &skill_.CreateRollbackResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill_.StandardizedError{})
	h.ResponseType(409, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Submit a target skill version to rollback to. Only one rollback or publish operation can be outstanding for a given skillId.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "defines the request body to create a rollback request",
   "in": "body",
   "name": "createRollbackRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.CreateRollbackRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Rollback request created; Returns the generated identifier to track the rollback request and returns a URL to track the status in Location header.",
   "headers": {
    "Content-Type": {
     "description": "contains content type of the response; only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Relative url to track rollback status",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.CreateRollbackResponse"
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
  "409": {
   "description": "The request could not be completed due to a conflict with the current state of the target resource.",
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
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "rollbackSkillV1"
}
*/
