package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	version_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/version"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateInteractionModelCatalogVersionV1 Create a new version of catalog entity for the given catalogId.

	catalogId - Provides a unique identifier of the catalog
	catalog -
*/
func (s *Client) CreateInteractionModelCatalogVersionV1(catalogId string, catalog *version_.VersionData) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs/{catalogId}/versions")
	h.Path("catalogId", catalogId)
	h.Body = catalog
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
 "description": "Create a new version of catalog entity for the given catalogId.\n",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
   "type": "string"
  },
  {
   "in": "body",
   "name": "catalog",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.VersionData"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Returns update status location link on success.",
   "headers": {
    "Location": {
     "description": "The location of the catalog status to track.",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. the catalog definition is invalid.",
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
   "description": "The specified catalog does not exist.",
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
 "x-operation-name": "createInteractionModelCatalogVersionV1"
}
*/
