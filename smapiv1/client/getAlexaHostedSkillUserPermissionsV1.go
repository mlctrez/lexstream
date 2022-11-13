package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	AlexaHosted_ "github.com/mlctrez/lexstream/smapiv1/skill/AlexaHosted"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetAlexaHostedSkillUserPermissionsV1 Get the current user permissions about Alexa hosted skill features.

	vendorId - vendorId
	permission - The permission of a hosted skill feature that customer needs to check.
*/
func (s *Client) GetAlexaHostedSkillUserPermissionsV1(vendorId string, permission string) (response *AlexaHosted_.HostedSkillPermission, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/vendors/{vendorId}/alexaHosted/permissions/{permission}")
	h.Path("vendorId", vendorId)
	h.Path("permission", permission)
	response = &AlexaHosted_.HostedSkillPermission{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the current user permissions about Alexa hosted skill features.",
 "parameters": [
  {
   "description": "vendorId",
   "in": "path",
   "name": "vendorId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The permission of a hosted skill feature that customer needs to check.",
   "in": "path",
   "name": "permission",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "response contains the user's permission of hosted skill features",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only \"application/json\" supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillPermission"
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
 "x-operation-name": "getAlexaHostedSkillUserPermissionsV1"
}
*/
