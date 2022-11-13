package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CloneLocaleV1 Creates a new clone locale workflow for a skill with given skillId, source locale, and target locales. In a single workflow, a locale can be cloned to multiple target locales. However, only one such workflow can be started at any time.

	skillId - The skill ID.
	stageV2 - Stages of a skill on which locales can be cloned. Currently only `development` stage is supported.

* `development` - skills which are currently in development corresponds to this stage.

	cloneLocaleRequest - Defines the request body for the cloneLocale API.
*/
func (s *Client) CloneLocaleV1(skillId string, stageV2 string, cloneLocaleRequest *skill_.CloneLocaleRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/cloneLocale")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Body = cloneLocaleRequest
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill_.StandardizedError{})
	h.ResponseType(409, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new clone locale workflow for a skill with given skillId, source locale, and target locales. In a single workflow, a locale can be cloned to multiple target locales. However, only one such workflow can be started at any time.\n",
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
   "description": "Defines the request body for the cloneLocale API.",
   "in": "body",
   "name": "cloneLocaleRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.CloneLocaleRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted.",
   "headers": {
    "Location": {
     "description": "Returns relative URL to track the progress of the workflow.",
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
  "409": {
   "description": "The request could not be completed due to a conflict with the current state of the target resource.",
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
 "x-operation-name": "cloneLocaleV1"
}
*/
