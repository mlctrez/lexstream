package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateSkillPackageV1 Creates a new import for a skill.

	createSkillWithPackageRequest - Defines the request body for createPackage API.
*/
func (s *Client) CreateSkillPackageV1(createSkillWithPackageRequest *skill_.CreateSkillWithPackageRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/imports")
	h.Body = createSkillWithPackageRequest
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(413, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new import for a skill.\n",
 "parameters": [
  {
   "description": "Defines the request body for createPackage API.",
   "in": "body",
   "name": "createSkillWithPackageRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.createSkillWithPackageRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted.",
   "headers": {
    "Location": {
     "description": "Contains relative URL to track import.",
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
  "413": {
   "description": "Payload too large.",
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
 "x-operation-name": "createSkillPackageV1"
}
*/
