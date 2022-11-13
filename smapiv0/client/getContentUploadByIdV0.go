package client

import (
	smapiv0 "github.com/mlctrez/lexstream/smapiv0"
	upload_ "github.com/mlctrez/lexstream/smapiv0/catalog/upload"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetContentUploadByIdV0 Gets detailed information about an upload which was created for a specific catalog. Includes the upload's ingestion steps and a presigned url for downloading the file.

	catalogId - Provides a unique identifier of the catalog
	uploadId - Unique identifier of the upload
*/
func (s *Client) GetContentUploadByIdV0(catalogId string, uploadId string) (response *upload_.GetContentUploadResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v0/catalogs/{catalogId}/uploads/{uploadId}")
	h.Path("catalogId", catalogId)
	h.Path("uploadId", uploadId)
	response = &upload_.GetContentUploadResponse{}
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
 "description": "Gets detailed information about an upload which was created for a specific catalog. Includes the upload's ingestion steps and a presigned url for downloading the file.",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Unique identifier of the upload",
   "in": "path",
   "name": "uploadId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successful operation.",
   "schema": {
    "$ref": "#/definitions/v0.catalog.upload.GetContentUploadResponse"
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
 "x-operation-name": "getContentUploadByIdV0"
}
*/
