package client

import (
	type1_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/type1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetInteractionModelSlotTypeDefinitionV1 Get the slot type definition.

	slotTypeId - The identifier for a slot type.
*/
func (s *Client) GetInteractionModelSlotTypeDefinitionV1(slotTypeId string) (response *type1_.SlotTypeDefinitionOutput, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}")
	h.Path("slotTypeId", slotTypeId)
	response = &type1_.SlotTypeDefinitionOutput{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the slot type definition.\n",
 "parameters": [
  {
   "description": "The identifier for a slot type.",
   "in": "path",
   "name": "slotTypeId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "The slot type definition.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeDefinitionOutput"
   }
  },
  "400": {
   "description": "The slot type cannot be retrieved due to errors listed.",
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
 "x-operation-name": "getInteractionModelSlotTypeDefinitionV1"
}
*/
