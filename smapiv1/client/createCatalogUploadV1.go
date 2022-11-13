package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	upload_ "github.com/mlctrez/lexstream/smapiv1/catalog/upload"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateCatalogUploadV1 Creates a new upload for a catalog and returns location to track the upload process.

	catalogId - Provides a unique identifier of the catalog
	catalogUploadRequestBody - Provides the request body for create content upload
*/
func (s *Client) CreateCatalogUploadV1(catalogId string, catalogUploadRequestBody *upload_.CatalogUploadBase) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/catalogs/{catalogId}/uploads")
	h.Path("catalogId", catalogId)
	h.Body = catalogUploadRequestBody
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
 "description": "Creates a new upload for a catalog and returns location to track the upload process.",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Provides the request body for create content upload",
   "in": "body",
   "name": "CatalogUploadRequestBody",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.catalog.upload.CatalogUploadBase"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted",
   "headers": {
    "Location": {
     "description": "Contains relative URL to track upload.",
     "type": "string"
    }
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
   "description": "Exceed the permitted request limit. Throttling criteria includes\ntotal requests, per API, ClientId, and CustomerId.\n",
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
 "summary": "Create new upload",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createCatalogUploadV1"
}
*/
