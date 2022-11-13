package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	catalog_ "github.com/mlctrez/lexstream/smapiv1/catalog"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GenerateCatalogUploadUrlV1 Generates preSigned urls to upload data

	catalogId - Provides a unique identifier of the catalog
	generateCatalogUploadUrlRequestBody - Request body to generate catalog upload url
*/
func (s *Client) GenerateCatalogUploadUrlV1(catalogId string, generateCatalogUploadUrlRequestBody *catalog_.CreateContentUploadUrlRequest) (response *catalog_.CreateContentUploadUrlResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/catalogs/{catalogId}/urls")
	h.Path("catalogId", catalogId)
	h.Body = generateCatalogUploadUrlRequestBody
	response = &catalog_.CreateContentUploadUrlResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Generates preSigned urls to upload data",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Request body to generate catalog upload url",
   "in": "body",
   "name": "GenerateCatalogUploadUrlRequestBody",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.catalog.CreateContentUploadUrlRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Successful operation.",
   "schema": {
    "$ref": "#/definitions/v1.catalog.CreateContentUploadUrlResponse"
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
     "description": "defines the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "generateCatalogUploadUrlV1"
}
*/
