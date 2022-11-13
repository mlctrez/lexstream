package client

import (
	version_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/version"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetInteractionModelCatalogVersionV1 Get catalog version data of given catalog version.

	catalogId - Provides a unique identifier of the catalog
	version - Version for interaction model.
*/
func (s *Client) GetInteractionModelCatalogVersionV1(catalogId string, version string) (response *version_.CatalogVersionData, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs/{catalogId}/versions/{version}")
	h.Path("catalogId", catalogId)
	h.Path("version", version)
	response = &version_.CatalogVersionData{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get catalog version data of given catalog version.\n",
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
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the catalog version metadata for the given catalogId and version.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.CatalogVersionData"
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
 "x-operation-name": "getInteractionModelCatalogVersionV1"
}
*/
