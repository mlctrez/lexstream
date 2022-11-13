package client

import (
	typeVersion_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/typeVersion"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetInteractionModelSlotTypeVersionV1 Get slot type version data of given slot type version.

	slotTypeId - The identifier for a slot type.
	version - Version for interaction model.
*/
func (s *Client) GetInteractionModelSlotTypeVersionV1(slotTypeId string, version string) (response *typeVersion_.SlotTypeVersionData, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}/versions/{version}")
	h.Path("slotTypeId", slotTypeId)
	h.Path("version", version)
	response = &typeVersion_.SlotTypeVersionData{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get slot type version data of given slot type version.\n",
 "parameters": [
  {
   "description": "The identifier for a slot type.",
   "in": "path",
   "name": "slotTypeId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Version for interaction model.",
   "in": "path",
   "name": "version",
   "pattern": "^~current|^~latest|[0-9]+$",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the slot type version metadata for the given slotTypeId and version.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.SlotTypeVersionData"
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
   "description": "There is no slot type defined for the slotTypeId.",
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
 "x-operation-name": "getInteractionModelSlotTypeVersionV1"
}
*/
