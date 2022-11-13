package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
StartBetaTestV1 Start a beta test for a given Alexa skill. System will send invitation emails to each tester in the test, and add entitlement on the acceptance.

	skillId - The skill ID.
*/
func (s *Client) StartBetaTestV1(skillId string) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/betaTest/start")
	h.Path("skillId", skillId)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Start a beta test for a given Alexa skill. System will send invitation emails to each tester in the test, and add entitlement on the acceptance.\n",
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
  "202": {
   "description": "Accept. Return a URL to track the resource in 'Location' header.",
   "headers": {
    "Content-Type": {
     "description": "Content type of the response; only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Relative URL to get the resource.",
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
  "409": {
   "description": "The request could not be completed due to a conflict with the current state of the target resource.",
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
 "summary": "Start beta test",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "startBetaTestV1"
}
*/
