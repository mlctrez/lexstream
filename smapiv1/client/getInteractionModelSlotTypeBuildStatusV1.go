package client

import (
	type1_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/type1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetInteractionModelSlotTypeBuildStatusV1 Get the status of slot type resource and its sub-resources for a given slotTypeId.

	slotTypeId - The identifier for a slot type.
	updateRequestId - The identifier for slotType version creation process
*/
func (s *Client) GetInteractionModelSlotTypeBuildStatusV1(slotTypeId string, updateRequestId string) (response *type1_.SlotTypeStatus, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/slotTypes/{slotTypeId}/updateRequest/{updateRequestId}")
	h.Path("slotTypeId", slotTypeId)
	h.Path("updateRequestId", updateRequestId)
	response = &type1_.SlotTypeStatus{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the status of slot type resource and its sub-resources for a given slotTypeId.\n",
 "parameters": [
  {
   "description": "The identifier for a slot type.",
   "in": "path",
   "name": "slotTypeId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The identifier for slotType version creation process",
   "in": "path",
   "name": "updateRequestId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the build status and error codes for the given slotTypeId.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeStatus"
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
 "x-operation-name": "getInteractionModelSlotTypeBuildStatusV1"
}
*/
