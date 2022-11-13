package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	catalog_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/catalog"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetInteractionModelCatalogUpdateStatusV1 Get the status of catalog resource and its sub-resources for a given catalogId.

	catalogId - Provides a unique identifier of the catalog
	updateRequestId - The identifier for slotType version creation process
*/
func (s *Client) GetInteractionModelCatalogUpdateStatusV1(catalogId string, updateRequestId string) (response *catalog_.CatalogStatus, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs/{catalogId}/updateRequest/{updateRequestId}")
	h.Path("catalogId", catalogId)
	h.Path("updateRequestId", updateRequestId)
	response = &catalog_.CatalogStatus{}
	h.Response = response
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
 "description": "Get the status of catalog resource and its sub-resources for a given catalogId.\n",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
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
   "description": "Returns the build status and error codes for the given catalogId",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogStatus"
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
   "description": "There is no catalog defined for the catalogId.",
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
 "x-operation-name": "getInteractionModelCatalogUpdateStatusV1"
}
*/
