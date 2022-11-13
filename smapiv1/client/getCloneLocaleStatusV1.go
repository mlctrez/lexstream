package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetCloneLocaleStatusV1 Returns the status of a clone locale workflow associated with the unique identifier of cloneLocaleRequestId.

	skillId - The skill ID.
	stageV2 - Stages of a skill on which locales can be cloned. Currently only `development` stage is supported.

* `development` - skills which are currently in development corresponds to this stage.

	cloneLocaleRequestId - Defines the identifier for a clone locale workflow.

If set to ~latest, request returns the status of the latest clone locale workflow.
*/
func (s *Client) GetCloneLocaleStatusV1(skillId string, stageV2 string, cloneLocaleRequestId string) (response *skill_.CloneLocaleStatusResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/cloneLocaleRequests/{cloneLocaleRequestId}")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Path("cloneLocaleRequestId", cloneLocaleRequestId)
	response = &skill_.CloneLocaleStatusResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Returns the status of a clone locale workflow associated with the unique identifier of cloneLocaleRequestId.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stages of a skill on which locales can be cloned. Currently only `development` stage is supported.\n* `development` - skills which are currently in development corresponds to this stage.\n",
   "in": "path",
   "name": "stageV2",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the identifier for a clone locale workflow.\nIf set to ~latest, request returns the status of the latest clone locale workflow.\n",
   "in": "path",
   "name": "cloneLocaleRequestId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "OK.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only \"application/json\" supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.CloneLocaleStatusResponse"
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getCloneLocaleStatusV1"
}
*/
