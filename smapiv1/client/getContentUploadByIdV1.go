package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	upload_ "github.com/mlctrez/lexstream/smapiv1/catalog/upload"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetContentUploadByIdV1 Gets detailed information about an upload which was created for a specific catalog. Includes the upload's ingestion steps and a url for downloading the file.

	catalogId - Provides a unique identifier of the catalog
	uploadId - Unique identifier of the upload
*/
func (s *Client) GetContentUploadByIdV1(catalogId string, uploadId string) (response *upload_.GetContentUploadResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/catalogs/{catalogId}/uploads/{uploadId}")
	h.Path("catalogId", catalogId)
	h.Path("uploadId", uploadId)
	response = &upload_.GetContentUploadResponse{}
	h.Response = response
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
 "description": "Gets detailed information about an upload which was created for a specific catalog. Includes the upload's ingestion steps and a url for downloading the file.",
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
    "$ref": "#/definitions/v1.catalog.upload.GetContentUploadResponse"
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
 "summary": "Get upload",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getContentUploadByIdV1"
}
*/
