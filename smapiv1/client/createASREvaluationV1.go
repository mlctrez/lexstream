package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateASREvaluationV1 This is an asynchronous API that starts an evaluation against the ASR model built by the skill's interaction model. The operation outputs an evaluationId which allows the retrieval of the current status of the operation and the results upon completion. This operation is unified, meaning both internal and external skill developers may use it to evaluate ASR models.

	postAsrEvaluationsRequest - Payload sent to trigger evaluation run.
	skillId - The skill ID.
*/
func (s *Client) CreateASREvaluationV1(postAsrEvaluationsRequest *evaluations_.PostAsrEvaluationsRequestObject, skillId string) (response *evaluations_.PostAsrEvaluationsResponseObject, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/asrEvaluations")
	h.Body = postAsrEvaluationsRequest
	h.Path("skillId", skillId)
	response = &evaluations_.PostAsrEvaluationsResponseObject{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(409, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an asynchronous API that starts an evaluation against the ASR model built by the skill's interaction model. The operation outputs an evaluationId which allows the retrieval of the current status of the operation and the results upon completion. This operation is unified, meaning both internal and external skill developers may use it to evaluate ASR models.\n",
 "parameters": [
  {
   "description": "Payload sent to trigger evaluation run.",
   "in": "body",
   "name": "PostAsrEvaluationsRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.PostAsrEvaluationsRequestObject"
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
     "description": "returned content type, - application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    },
    "Location": {
     "description": "URI indicating where to poll the status of the evaluation.\n",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.evaluations.PostAsrEvaluationsResponseObject"
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
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "default": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Start an evaluation against the ASR model built by the skill's interaction model.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createASREvaluationV1"
}
*/
