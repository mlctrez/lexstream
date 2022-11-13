package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetNLUEvaluationV1 API which requests top level information about the evaluation like the current state of the job, status of the evaluation (if complete). Also returns data used to start the job, like the number of test cases, stage, locale, and start time. This should be considered the 'cheap' operation while getResultForNLUEvaluations is 'expensive'.

	skillId - The skill ID.
	evaluationId - Identifier of the evaluation.
*/
func (s *Client) GetNLUEvaluationV1(skillId string, evaluationId string) (response *evaluations_.GetNLUEvaluationResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/nluEvaluations/{evaluationId}")
	h.Path("skillId", skillId)
	h.Path("evaluationId", evaluationId)
	response = &evaluations_.GetNLUEvaluationResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which requests top level information about the evaluation like the current state of the job, status of the evaluation (if complete). Also returns data used to start the job, like the number of test cases, stage, locale, and start time. This should be considered the 'cheap' operation while getResultForNLUEvaluations is 'expensive'.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the evaluation.",
   "in": "path",
   "name": "evaluationId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Evaluation exists and its status is queryable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, - application/json or application/hal+json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.GetNLUEvaluationResponse"
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
 "summary": "Get top level information and status of a nlu evaluation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getNLUEvaluationV1"
}
*/
