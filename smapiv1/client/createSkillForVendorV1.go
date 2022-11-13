package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateSkillForVendorV1 Creates a new skill for given vendorId.

	createSkillRequest - Defines the request body for createSkill API.
*/
func (s *Client) CreateSkillForVendorV1(createSkillRequest *skill_.CreateSkillRequest) (response *skill_.CreateSkillResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills")
	h.Body = createSkillRequest
	response = &skill_.CreateSkillResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new skill for given vendorId.",
 "parameters": [
  {
   "description": "Defines the request body for createSkill API.",
   "in": "body",
   "name": "createSkillRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.createSkillRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted; Returns a URL to track the status in 'Location' header.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Contains relative URL to get skill status.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.CreateSkillResponse"
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
 "x-operation-name": "createSkillForVendorV1"
}
*/
