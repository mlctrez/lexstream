package client

import (
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateNLUEvaluationsV1 This is an asynchronous API that starts an evaluation against the NLU model built by the skill's interaction model.
The operation outputs an evaluationId which allows the retrieval of the current status of the operation and the results upon completion. This operation is unified, meaning both internal and external skill developers may use it evaluate NLU models.

	evaluateNLURequest - Payload sent to the evaluate NLU API.
	skillId - The skill ID.
*/
func (s *Client) CreateNLUEvaluationsV1(evaluateNLURequest *evaluations_.EvaluateNLURequest, skillId string) (response *evaluations_.EvaluateResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/nluEvaluations")
	h.Body = evaluateNLURequest
	h.Path("skillId", skillId)
	response = &evaluations_.EvaluateResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an asynchronous API that starts an evaluation against the NLU model built by the skill's interaction model.\nThe operation outputs an evaluationId which allows the retrieval of the current status of the operation and the results upon completion. This operation is unified, meaning both internal and external skill developers may use it evaluate NLU models.\n",
 "parameters": [
  {
   "description": "Payload sent to the evaluate NLU API.",
   "in": "body",
   "name": "evaluateNLURequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.EvaluateNLURequest"
   }
  },
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
   "description": "Evaluation has successfully begun.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, - application/json or application/hal+json supported",
     "type": "string"
    },
    "Location": {
     "description": "GET Location of the started evaluation.\n",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.EvaluateResponse"
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
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
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
 "summary": "Start an evaluation against the NLU model built by the skill's interaction model.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createNLUEvaluationsV1"
}
*/
