package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	catalog_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/catalog"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateInteractionModelCatalogV1 Create a new version of catalog within the given catalogId.

	catalog -
*/
func (s *Client) CreateInteractionModelCatalogV1(catalog *catalog_.DefinitionData) (response *catalog_.CatalogResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs")
	h.Body = catalog
	response = &catalog_.CatalogResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(412, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Create a new version of catalog within the given catalogId.\n",
 "parameters": [
  {
   "in": "body",
   "name": "catalog",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.catalog.DefinitionData"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Returns the generated catalogId.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogResponse"
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
  "412": {
   "description": "Precondition failed.",
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
 "x-operation-name": "createInteractionModelCatalogV1"
}
*/
