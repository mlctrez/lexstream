package client

import (
	type1_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/type1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateInteractionModelSlotTypeV1 Update description and vendorGuidance string for certain version of a slot type.

	slotTypeId - The identifier for a slot type.
	updateRequest -
*/
func (s *Client) UpdateInteractionModelSlotTypeV1(slotTypeId string, updateRequest *type1_.UpdateRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}/update")
	h.Path("slotTypeId", slotTypeId)
	h.Body = updateRequest
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
   "in": "body",
   "name": "updateRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.UpdateRequest"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "No content, indicates the fields were successfully updated."
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
 "x-operation-name": "updateInteractionModelSlotTypeV1"
}
*/
