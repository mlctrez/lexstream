package client

import (
	smartHomeEvaluation_ "github.com/mlctrez/lexstream/smapiv1/smartHomeEvaluation"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSmartHomeCapabilityEvaluationV1 Get top level information and status of a Smart Home capability evaluation.

	skillId - The skill ID.
	evaluationId - A unique ID to identify each Smart Home capability evaluation.
*/
func (s *Client) GetSmartHomeCapabilityEvaluationV1(skillId string, evaluationId string) (response *smartHomeEvaluation_.GetSHCapabilityEvaluationResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/smartHome/testing/capabilityEvaluations/{evaluationId}")
	h.Path("skillId", skillId)
	h.Path("evaluationId", evaluationId)
	response = &smartHomeEvaluation_.GetSHCapabilityEvaluationResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get top level information and status of a Smart Home capability evaluation.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "A unique ID to identify each Smart Home capability evaluation.",
   "in": "path",
   "name": "evaluationId",
   "required": true,
   "type": "string"
  }
 ],
 "produces": [
  "application/json"
 ],
 "responses": {
  "200": {
   "description": "Successfully retrieved the evaluation status.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.GetSHCapabilityEvaluationResponse"
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
 "summary": "Get top level information and status of a Smart Home capability evaluation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getSmartHomeCapabilityEvaluationV1"
}
*/
