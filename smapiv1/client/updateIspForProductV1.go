package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	isp_ "github.com/mlctrez/lexstream/smapiv1/isp"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateIspForProductV1 Updates in-skill product definition for given productId. Only development stage supported.

	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
	productId - The in-skill product ID.
	stage - Stage for skill.
	updateInSkillProductRequest - defines the request body for updateInSkillProduct API.
*/
func (s *Client) UpdateIspForProductV1(if_Match string, productId string, stage string, updateInSkillProductRequest *isp_.UpdateInSkillProductRequest) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/inSkillProducts/{productId}/stages/{stage}")
	h.Header("If-Match", if_Match)
	h.Path("productId", productId)
	h.Path("stage", stage)
	h.Body = updateInSkillProductRequest
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(412, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Updates in-skill product definition for given productId. Only development stage supported.",
 "parameters": [
  {
   "description": "Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.",
   "in": "header",
   "name": "If-Match",
   "required": false,
   "type": "string"
  },
  {
   "description": "The in-skill product ID.",
   "in": "path",
   "name": "productId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage for skill.",
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "defines the request body for updateInSkillProduct API.",
   "in": "body",
   "name": "updateInSkillProductRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.isp.updateInSkillProductRequest"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "Success.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Bad request. Returned when a required parameter is not present, badly formatted.\n",
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
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "Request is forbidden.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "Requested resource not found.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "412": {
   "description": "Precondition failed.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Too many requests received.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "updateIspForProductV1"
}
*/
