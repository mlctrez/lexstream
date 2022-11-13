package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	accountLinking_ "github.com/mlctrez/lexstream/smapiv1/skill/accountLinking"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateAccountLinkingInfoV1 Create AccountLinking information for the skill.

	skillId - The skill ID.
	stageV2 - Stages of a skill including the new certified stage.

* `development` - skills which are currently in development corresponds to this stage.
* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.
* `live` - skills which are currently live corresponds to this stage.

	accountLinkingRequest - The fields required to create accountLinking partner.
	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
*/
func (s *Client) UpdateAccountLinkingInfoV1(skillId string, stageV2 string, accountLinkingRequest *accountLinking_.AccountLinkingRequest, if_Match string) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/accountLinkingClient")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Body = accountLinkingRequest
	h.Header("If-Match", if_Match)
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(412, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Create AccountLinking information for the skill.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stages of a skill including the new certified stage.\n* `development` - skills which are currently in development corresponds to this stage.\n* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.\n* `live` - skills which are currently live corresponds to this stage.\n",
   "in": "path",
   "name": "stageV2",
   "required": true,
   "type": "string"
  },
  {
   "description": "The fields required to create accountLinking partner.",
   "in": "body",
   "name": "accountLinkingRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingRequest"
   }
  },
  {
   "description": "Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.",
   "in": "header",
   "name": "If-Match",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "Success",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only 'application/json' supported.",
     "type": "string"
    },
    "ETag": {
     "description": "Identifer for the updated version of the resource, can be used for conditional updates.",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. Authorization Url is invalid.",
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
  "412": {
   "description": "Precondition failed.",
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
 "x-operation-name": "updateAccountLinkingInfoV1"
}
*/
