package client

import (
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ImportSkillPackageV1 Creates a new import for a skill with given skillId.

	updateSkillWithPackageRequest - Defines the request body for updatePackage API.
	skillId - The skill ID.
	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
*/
func (s *Client) ImportSkillPackageV1(updateSkillWithPackageRequest *skill_.UpdateSkillWithPackageRequest, skillId string, if_Match string) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/imports")
	h.Body = updateSkillWithPackageRequest
	h.Path("skillId", skillId)
	h.Header("If-Match", if_Match)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new import for a skill with given skillId.\n",
 "parameters": [
  {
   "description": "Defines the request body for updatePackage API.",
   "in": "body",
   "name": "updateSkillWithPackageRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.updateSkillWithPackageRequest"
   }
  },
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
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
   "description": "Accepted.",
   "headers": {
    "Location": {
     "description": "Contains relative URL to track import.",
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
  "413": {
   "description": "Payload too large.",
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
 "x-operation-name": "importSkillPackageV1"
}
*/
