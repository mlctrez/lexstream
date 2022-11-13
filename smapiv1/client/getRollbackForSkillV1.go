package client

import (
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetRollbackForSkillV1 Get the rollback status of a skill given an associated rollbackRequestId. Use ~latest in place of rollbackRequestId to get the latest rollback status.

	skillId - The skill ID.
	rollbackRequestId - Defines the identifier for a rollback request. If set to ~latest, request returns the status of the latest rollback request.
*/
func (s *Client) GetRollbackForSkillV1(skillId string, rollbackRequestId string) (response *skill_.RollbackRequestStatus, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/rollbacks/{rollbackRequestId}")
	h.Path("skillId", skillId)
	h.Path("rollbackRequestId", rollbackRequestId)
	response = &skill_.RollbackRequestStatus{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the rollback status of a skill given an associated rollbackRequestId. Use ~latest in place of rollbackRequestId to get the latest rollback status.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the identifier for a rollback request. If set to ~latest, request returns the status of the latest rollback request.",
   "in": "path",
   "name": "rollbackRequestId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the rollback status for a given skillId and rollbackRequestId. Returns the latest rollback status if ~latest is used in place of rollbackRequestId.",
   "headers": {
    "Content-Type": {
     "description": "contains content type of the response; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.RollbackRequestStatus"
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
 "x-operation-name": "getRollbackForSkillV1"
}
*/
