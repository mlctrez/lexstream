package client

import (
	Manifest_ "github.com/mlctrez/lexstream/smapiv1/skill/Manifest"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateSkillManifestV1 Updates skill manifest for given skillId and stage.

	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
	skillId - The skill ID.
	stageV2 - Stages of a skill including the new certified stage.

* `development` - skills which are currently in development corresponds to this stage.
* `certified` -  skills which have completed certification and ready for publishing corresponds to this stage.
* `live` - skills which are currently live corresponds to this stage.

	updateSkillRequest - Defines the request body for updateSkill API.
*/
func (s *Client) UpdateSkillManifestV1(if_Match string, skillId string, stageV2 string, updateSkillRequest *Manifest_.SkillManifestEnvelope) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/stages/{stageV2}/manifest")
	h.Header("If-Match", if_Match)
	h.Path("skillId", skillId)
	h.Path("stageV2", stageV2)
	h.Body = updateSkillRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Updates skill manifest for given skillId and stage.",
 "parameters": [
  {
   "description": "Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.",
   "in": "header",
   "name": "If-Match",
   "required": false,
   "type": "string"
  },
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
   "description": "Defines the request body for updateSkill API.",
   "in": "body",
   "name": "updateSkillRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEnvelope"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted; Returns a URL to track the status in 'Location' header.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Contains relative URL to get skill status.",
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
  "412": {
   "description": "Precondition failed.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
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
 "x-operation-name": "updateSkillManifestV1"
}
*/
