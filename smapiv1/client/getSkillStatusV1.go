package client

import (
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSkillStatusV1 Get the status of skill resource and its sub-resources for a given skillId.

	skillId - The skill ID.
	resource - Resource name for which status information is desired.

It is an optional, filtering parameter and can be used more than once, to retrieve status for all the desired (sub)resources only, in single API call.
If this parameter is not specified, status for all the resources/sub-resources will be returned.
*/
func (s *Client) GetSkillStatusV1(skillId string, resource string) (response *skill_.SkillStatus, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/status")
	h.Path("skillId", skillId)
	h.Param("resource", resource)
	response = &skill_.SkillStatus{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the status of skill resource and its sub-resources for a given skillId.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Resource name for which status information is desired.\nIt is an optional, filtering parameter and can be used more than once, to retrieve status for all the desired (sub)resources only, in single API call.\nIf this parameter is not specified, status for all the resources/sub-resources will be returned.\n",
   "in": "query",
   "name": "resource",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns status for skill resource and sub-resources.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.SkillStatus"
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
 "x-operation-name": "getSkillStatusV1"
}
*/
