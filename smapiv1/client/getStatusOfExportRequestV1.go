package client

import (
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetStatusOfExportRequestV1 Get status for given exportId

	exportId - The Export ID.
*/
func (s *Client) GetStatusOfExportRequestV1(exportId string) (response *skill_.ExportResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/exports/{exportId}")
	h.Path("exportId", exportId)
	response = &skill_.ExportResponse{}
	h.Response = response
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(404, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get status for given exportId\n",
 "parameters": [
  {
   "description": "The Export ID.",
   "in": "path",
   "name": "exportId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "OK.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.ExportResponse"
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
  "404": {
   "description": "The resource being requested is not found.",
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
 "x-operation-name": "getStatusOfExportRequestV1"
}
*/
