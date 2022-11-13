package client

import (
	type1_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/type1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateInteractionModelSlotTypeV1 Create a new version of slot type within the given slotTypeId.

	slotType -
*/
func (s *Client) CreateInteractionModelSlotTypeV1(slotType *type1_.DefinitionData) (response *type1_.SlotTypeResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes")
	h.Body = slotType
	response = &type1_.SlotTypeResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Create a new version of slot type within the given slotTypeId.\n",
 "parameters": [
  {
   "in": "body",
   "name": "slotType",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.DefinitionData"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the generated slotTypeId.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeResponse"
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
 "x-operation-name": "createInteractionModelSlotTypeV1"
}
*/
