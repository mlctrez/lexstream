package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	upload_ "github.com/mlctrez/lexstream/smapiv0/catalog/upload"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateContentUploadV0 Creates a new upload for a catalog and returns presigned upload parts for uploading the file.

	catalogId - Provides a unique identifier of the catalog
	createContentUploadRequest - Defines the request body for updateCatalog API.
*/
func (s *Client) CreateContentUploadV0(catalogId string, createContentUploadRequest *upload_.CreateContentUploadRequest) (response *upload_.CreateContentUploadResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v0/catalogs/{catalogId}/uploads")
	h.Path("catalogId", catalogId)
	h.Body = createContentUploadRequest
	response = &upload_.CreateContentUploadResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv0.BadRequestError{})
	h.ResponseType(401, &smapiv0.Error{})
	h.ResponseType(403, &smapiv0.BadRequestError{})
	h.ResponseType(404, &smapiv0.Error{})
	h.ResponseType(429, &smapiv0.Error{})
	h.ResponseType(500, &smapiv0.Error{})
	h.ResponseType(503, &smapiv0.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new upload for a catalog and returns presigned upload parts for uploading the file.",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the request body for updateCatalog API.",
   "in": "body",
   "name": "CreateContentUploadRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v0.catalog.upload.CreateContentUploadRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Content upload created.",
   "schema": {
    "$ref": "#/definitions/v0.catalog.upload.CreateContentUploadResponse"
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
 "x-operation-name": "createContentUploadV0"
}
*/
