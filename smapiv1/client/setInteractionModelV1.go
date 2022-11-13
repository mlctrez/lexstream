package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	interactionModel_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetInteractionModelV1 Creates an `InteractionModel` for the skill.

	skillId - The skill ID.
	stageV2 - Stages of a skill including the new certified stage.

* `development` - skills which are currently in development corresponds to this stage.
* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.
* `live` - skills which are currently live corresponds to this stage.

	locale - The locale for the model requested e.g. en-GB, en-US, de-DE.
	interactionModel -
	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
*/
func (s *Client) SetInteractionModelV1(skillId string, stageV2 string, locale string, interactionModel *interactionModel_.InteractionModelData, if_Match string) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/interactionModel/locales/{locale}")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Path("locale", locale)
	h.Body = interactionModel
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
 "description": "Creates an `InteractionModel` for the skill.\n",
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
   "description": "The locale for the model requested e.g. en-GB, en-US, de-DE.",
   "format": "languager-region; same as BCP-47 language tag format",
   "in": "path",
   "name": "locale",
   "required": true,
   "type": "string"
  },
  {
   "in": "body",
   "name": "interactionModel",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.InteractionModelData"
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
  "202": {
   "description": "Returns build status location link on success.",
   "headers": {
    "Location": {
     "description": "Build status URL.",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. the input interaction model is invalid.",
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
   "description": "The specified skill or stage or locale does not exist.",
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
 "x-operation-name": "setInteractionModelV1"
}
*/
