package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
DeleteSkillEnablementV1 Deletes the enablement for given skillId/stage and customerId (retrieved from Auth token).

	skillId - The skill ID.
	stage - Stage for skill.
*/
func (s *Client) DeleteSkillEnablementV1(skillId string, stage string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/enablement")
	h.Path("skillId", skillId)
	h.Path("stage", stage)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Deletes the enablement for given skillId/stage and customerId (retrieved from Auth token).",
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
  }
 ],
 "responses": {
  "204": {
   "description": "No Content; Confirms that enablement is successfully deleted.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
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
 "x-operation-name": "deleteSkillEnablementV1"
}
*/
