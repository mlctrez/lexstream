package client

import (
	accountLinking_ "github.com/mlctrez/lexstream/smapiv1/skill/accountLinking"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetAccountLinkingInfoV1 Get AccountLinking information for the skill.

	skillId - The skill ID.
	stageV2 - Stages of a skill including the new certified stage.

* `development` - skills which are currently in development corresponds to this stage.
* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.
* `live` - skills which are currently live corresponds to this stage.
*/
func (s *Client) GetAccountLinkingInfoV1(skillId string, stageV2 string) (response *accountLinking_.AccountLinkingResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/accountLinkingClient")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	response = &accountLinking_.AccountLinkingResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get AccountLinking information for the skill.\n",
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
  }
 ],
 "responses": {
  "200": {
   "description": "Returns AccountLinking response of the skill.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    },
    "ETag": {
     "description": "Identifer for the version of the resource, can be used for conditional updates.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingResponse"
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
 "x-operation-name": "getAccountLinkingInfoV1"
}
*/
