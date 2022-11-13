package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
DeleteASRAnnotationSetV1 API which deletes the ASR annotation set. Developers cannot get/list the deleted annotation set.

	skillId - The skill ID.
	annotationSetId - Identifier of the ASR annotation set.
*/
func (s *Client) DeleteASRAnnotationSetV1(skillId string, annotationSetId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets/{annotationSetId}")
	h.Path("skillId", skillId)
	h.Path("annotationSetId", annotationSetId)
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(409, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which deletes the ASR annotation set. Developers cannot get/list the deleted annotation set.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the ASR annotation set.",
   "in": "path",
   "name": "annotationSetId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "ASR annotation set exists and is deleted successfully."
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
     "description": "Define the authentication method that should be used to gain access to a resource.",
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
  "409": {
   "description": "The request could not be completed due to a conflict with the current state of the target resource.",
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
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "default": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Delete the ASR annotation set",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "deleteASRAnnotationSetV1"
}
*/
