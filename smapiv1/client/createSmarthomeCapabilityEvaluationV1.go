package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	smartHomeEvaluation_ "github.com/mlctrez/lexstream/smapiv1/smartHomeEvaluation"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateSmarthomeCapabilityEvaluationV1 Start a capability evaluation against a Smart Home skill.

	skillId - The skill ID.
	evaluateSHCapabilityPayload - Payload sent to start a capability evaluation against a Smart Home skill.
*/
func (s *Client) CreateSmarthomeCapabilityEvaluationV1(skillId string, evaluateSHCapabilityPayload *smartHomeEvaluation_.EvaluateSHCapabilityRequest) (response *smartHomeEvaluation_.EvaluateSHCapabilityResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/smartHome/testing/capabilityEvaluations")
	h.Path("skillId", skillId)
	h.Body = evaluateSHCapabilityPayload
	response = &smartHomeEvaluation_.EvaluateSHCapabilityResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.BadRequestError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.BadRequestError{})
	h.ResponseType(409, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.BadRequestError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "consumes": [
  "application/json"
 ],
 "description": "Start a capability evaluation against a Smart Home skill.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to start a capability evaluation against a Smart Home skill.",
   "in": "body",
   "name": "EvaluateSHCapabilityPayload",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluateSHCapabilityRequest"
   }
  }
 ],
 "produces": [
  "application/json"
 ],
 "responses": {
  "200": {
   "description": "Evaluation has successfully begun.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "type": "string"
    },
    "Location": {
     "description": "Get location of the capability evaluaiton.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.EvaluateSHCapabilityResponse"
   }
  },
  "400": {
   "description": "Bad Request. Returned when the request payload is malformed or when, at least, one required property is missing or invalid.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "403": {
   "description": "API user does not have permission or is currently in a state that does not allow calls to this API.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill, test plan, or evaluation does not exist. The error response will contain a description that indicates the specific resource type that was not found.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "409": {
   "description": "A test run is already in progress for the specified endpoint. Please retry after some time.\n",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceeded the permitted request limit. Throttling criteria includes total requests, per API and CustomerId.\n",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "default": {
   "description": "Internal server error.\n",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Start a capability evaluation against a Smart Home skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createSmarthomeCapabilityEvaluationV1"
}
*/
