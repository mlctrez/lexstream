package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	typeVersion_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/typeVersion"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateInteractionModelSlotTypeVersionV1 Update description and vendorGuidance string for certain version of a slot type.

	slotTypeId - The identifier for a slot type.
	version - Version for interaction model.
	slotTypeUpdate -
*/
func (s *Client) UpdateInteractionModelSlotTypeVersionV1(slotTypeId string, version string, slotTypeUpdate *typeVersion_.SlotTypeUpdate) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}/versions/{version}/update")
	h.Path("slotTypeId", slotTypeId)
	h.Path("version", version)
	h.Body = slotTypeUpdate
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
 "description": "Update description and vendorGuidance string for certain version of a slot type.\n",
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
  },
  {
   "in": "body",
   "name": "slotTypeUpdate",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.slotTypeUpdate"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "No Content; Confirms that version is successfully updated."
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
     "description": "defines the authentication method that should be used to gain access to a resource.",
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
 "x-operation-name": "updateInteractionModelSlotTypeVersionV1"
}
*/
