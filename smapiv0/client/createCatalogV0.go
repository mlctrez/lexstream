package client

import (
	catalog_ "github.com/mlctrez/lexstream/smapiv0/catalog"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateCatalogV0 Creates a new catalog based on information provided in the request.

	createCatalogRequest - Defines the request body for createCatalog API.
*/
func (s *Client) CreateCatalogV0(createCatalogRequest *catalog_.CreateCatalogRequest) (response *catalog_.CatalogDetails, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v0/catalogs")
	h.Body = createCatalogRequest
	response = &catalog_.CatalogDetails{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new catalog based on information provided in the request.",
 "parameters": [
  {
   "description": "Defines the request body for createCatalog API.",
   "in": "body",
   "name": "CreateCatalogRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v0.catalog.CreateCatalogRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Catalog created.",
   "schema": {
    "$ref": "#/definitions/v0.catalog.CatalogDetails"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v0.BadRequestError"
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
    "$ref": "#/definitions/v0.Error"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v0.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v0.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createCatalogV0"
}
*/
