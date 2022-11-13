package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	betaTest_ "github.com/mlctrez/lexstream/smapiv1/skill/betaTest"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateBetaTestV1 Update a beta test for a given Alexa skill.

	skillId - The skill ID.
	createTestBody - JSON object containing the details of a beta test used to create the test.
*/
func (s *Client) UpdateBetaTestV1(skillId string, createTestBody *betaTest_.TestBody) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/betaTest")
	h.Path("skillId", skillId)
	h.Body = createTestBody
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Update a beta test for a given Alexa skill.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "JSON object containing the details of a beta test used to create the test.",
   "in": "body",
   "name": "createTestBody",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v1.skill.betaTest.TestBody"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "Success. No content.",
   "headers": {
    "Content-Type": {
     "description": "Content type of the response; only application/json supported.",
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
     "description": "Define the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Update beta test.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "updateBetaTestV1"
}
*/
