package client

import (
	simulations_ "github.com/mlctrez/lexstream/smapiv1/skill/simulations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SimulateSkillV1 This is an asynchronous API that simulates a skill execution in the Alexa eco-system given an utterance text of what a customer would say to Alexa. A successful response will contain a header with the location of the simulation resource. In cases where requests to this API results in an error, the response will contain an error code and a description of the problem. The skill being simulated must be in development stage, and it must also belong to and be enabled by the user of this API. Concurrent requests per user is currently not supported.

	skillId - The skill ID.
	simulationsApiRequest - Payload sent to the skill simulation API.
*/
func (s *Client) SimulateSkillV1(skillId string, simulationsApiRequest *simulations_.SimulationsApiRequest) (response *simulations_.SimulationsApiResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/simulations")
	h.Path("skillId", skillId)
	h.Body = simulationsApiRequest
	response = &simulations_.SimulationsApiResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an asynchronous API that simulates a skill execution in the Alexa eco-system given an utterance text of what a customer would say to Alexa. A successful response will contain a header with the location of the simulation resource. In cases where requests to this API results in an error, the response will contain an error code and a description of the problem. The skill being simulated must be in development stage, and it must also belong to and be enabled by the user of this API. Concurrent requests per user is currently not supported.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the skill simulation API.",
   "in": "body",
   "name": "simulationsApiRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.simulations.SimulationsApiRequest"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Skill simulation has successfully began.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Path to simulation resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.simulations.SimulationsApiResponse"
   }
  },
  "400": {
   "description": "Bad request due to invalid or missing data.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "API user does not have permission to call this API or is currently in a state that does not allow simulation of this skill.\n",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill does not exist.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "409": {
   "description": "This requests conflicts with another one currently being processed.\n",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "API user has exceeded the permitted request rate.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal service error.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Simulate executing a skill with the given id.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "simulateSkillV1"
}
*/
