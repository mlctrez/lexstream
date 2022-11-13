package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	typeVersion_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/typeVersion"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateInteractionModelSlotTypeVersionV1 Create a new version of slot type entity for the given slotTypeId.

	slotTypeId - The identifier for a slot type.
	slotType -
*/
func (s *Client) CreateInteractionModelSlotTypeVersionV1(slotTypeId string, slotType *typeVersion_.VersionData) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}/versions")
	h.Path("slotTypeId", slotTypeId)
	h.Body = slotType
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Create a new version of slot type entity for the given slotTypeId.\n",
 "parameters": [
  {
   "description": "The identifier for a slot type.",
   "in": "path",
   "name": "slotTypeId",
   "required": true,
   "type": "string"
  },
  {
   "in": "body",
   "name": "slotType",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.VersionData"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Returns update status location link on success.",
   "headers": {
    "Location": {
     "description": "The location of the slot type status to track.",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. the slot type definition is invalid.",
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
   "description": "The specified slot type does not exist.",
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
 "x-operation-name": "createInteractionModelSlotTypeVersionV1"
}
*/
