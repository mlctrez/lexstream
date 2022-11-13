package client

import (
	betaTest_ "github.com/mlctrez/lexstream/smapiv1/skill/betaTest"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetBetaTestV1 Get beta test for a given Alexa skill.

	skillId - The skill ID.
*/
func (s *Client) GetBetaTestV1(skillId string) (response *betaTest_.BetaTest, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/betaTest")
	h.Path("skillId", skillId)
	response = &betaTest_.BetaTest{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get beta test for a given Alexa skill.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Success.",
   "headers": {
    "Content-Type": {
     "description": "Content type of the response; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.betaTest.BetaTest"
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
 "summary": "Get beta test.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getBetaTestV1"
}
*/
