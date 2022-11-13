package client

import (
	testers_ "github.com/mlctrez/lexstream/smapiv1/skill/betaTest/testers"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
AddTestersToBetaTestV1 Add testers to a beta test for the given Alexa skill.  System will send invitation email to each tester and add entitlement on the acceptance.

	skillId - The skill ID.
	testersRequest - JSON object containing the email address of beta testers.
*/
func (s *Client) AddTestersToBetaTestV1(skillId string, testersRequest *testers_.TestersList) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/betaTest/testers/add")
	h.Path("skillId", skillId)
	h.Body = testersRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Add testers to a beta test for the given Alexa skill.  System will send invitation email to each tester and add entitlement on the acceptance.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "JSON object containing the email address of beta testers.",
   "in": "body",
   "name": "TestersRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.betaTest.testers.TestersList"
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
     "description": "Defines the authentication method that should be used to gain access to a resource.",
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
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
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
 "summary": "Add testers to an existing beta test.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "addTestersToBetaTestV1"
}
*/
