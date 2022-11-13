package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
GetInteractionModelMetadataV1 Get the latest metadata for the interaction model resource for the given stage.

	skillId - The skill ID.
	stageV2 - Stages of a skill including the new certified stage.

* `development` - skills which are currently in development corresponds to this stage.
* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.
* `live` - skills which are currently live corresponds to this stage.

	locale - The locale for the model requested e.g. en-GB, en-US, de-DE.
*/
func (s *Client) GetInteractionModelMetadataV1(skillId string, stageV2 string, locale string) (err error) {
	h := swaggerlt.NewRequestHelper("head", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/interactionModel/locales/{locale}")
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Path("locale", locale)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the latest metadata for the interaction model resource for the given stage.\n",
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
  }
 ],
 "responses": {
  "204": {
   "description": "Success. There is no content but returns etag.",
   "headers": {
    "ETag": {
     "description": "Identifer for the version of the resource, can be used for conditional updates",
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
   "description": "The specified skill or stage or locale does not exist",
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
 "x-operation-name": "getInteractionModelMetadataV1"
}
*/
