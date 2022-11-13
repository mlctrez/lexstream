package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	invocations_ "github.com/mlctrez/lexstream/smapiv1/skill/invocations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
InvokeSkillV1 This is a synchronous API that invokes the Lambda or third party HTTPS endpoint for a given skill. A successful response will contain information related to what endpoint was called, payload sent to and received from the endpoint. In cases where requests to this API results in an error, the response will contain an error code and a description of the problem. In cases where invoking the skill endpoint specifically fails, the response will contain a status attribute indicating that a failure occurred and details about what was sent to the endpoint. The skill must belong to and be enabled by the user of this API. Also, note that calls to the skill endpoint will timeout after 10 seconds.

	skillId - The skill ID.
	invokeSkillRequest - Payload sent to the skill invocation API.
*/
func (s *Client) InvokeSkillV1(skillId string, invokeSkillRequest *invocations_.InvokeSkillRequest) (response *invocations_.InvokeSkillResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/invocations")
	h.Path("skillId", skillId)
	h.Body = invokeSkillRequest
	response = &invocations_.InvokeSkillResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is a synchronous API that invokes the Lambda or third party HTTPS endpoint for a given skill. A successful response will contain information related to what endpoint was called, payload sent to and received from the endpoint. In cases where requests to this API results in an error, the response will contain an error code and a description of the problem. In cases where invoking the skill endpoint specifically fails, the response will contain a status attribute indicating that a failure occurred and details about what was sent to the endpoint. The skill must belong to and be enabled by the user of this API. Also, note that calls to the skill endpoint will timeout after 10 seconds.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the skill invocation API.",
   "in": "body",
   "name": "invokeSkillRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.invocations.InvokeSkillRequest"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Skill was invoked.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.skill.invocations.InvokeSkillResponse"
   }
  },
  "400": {
   "description": "Bad request due to invalid or missing data.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "403": {
   "description": "API user does not have permission to call this API or is currently in a state that does not allow invocation of this skill.\n",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill does not exist.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "API user has exceeded the permitted request rate.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "invokeSkillV1"
}
*/
