package client

import (
	version_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/version"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateInteractionModelCatalogVersionV1 Update description and vendorGuidance string for certain version of a catalog.

	catalogId - Provides a unique identifier of the catalog
	version - Version for interaction model.
	catalogUpdate -
*/
func (s *Client) UpdateInteractionModelCatalogVersionV1(catalogId string, version string, catalogUpdate *version_.CatalogUpdate) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs/{catalogId}/versions/{version}/update")
	h.Path("catalogId", catalogId)
	h.Path("version", version)
	h.Body = catalogUpdate
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Update description and vendorGuidance string for certain version of a catalog.\n",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
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
   "name": "catalogUpdate",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.catalogUpdate"
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
   "description": "There is no catalog defined for the catalogId",
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
 "x-operation-name": "updateInteractionModelCatalogVersionV1"
}
*/
