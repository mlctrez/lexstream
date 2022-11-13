package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
DeletePropertiesForNLUAnnotationSetsV1 API which deletes the NLU annotation set. Developers cannot get/list the deleted annotation set.

	skillId - The skill ID.
	annotationId - Identifier of the NLU annotation set.
*/
func (s *Client) DeletePropertiesForNLUAnnotationSetsV1(skillId string, annotationId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets/{annotationId}")
	h.Path("skillId", skillId)
	h.Path("annotationId", annotationId)
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which deletes the NLU annotation set. Developers cannot get/list the deleted annotation set.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the NLU annotation set.",
   "in": "path",
   "name": "annotationId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "NLU annotation set exists and is deleted successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/hal+json or application/json supported",
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
  }
 },
 "summary": "Delete the NLU annotation set",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "deletePropertiesForNLUAnnotationSetsV1"
}
*/
