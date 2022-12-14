package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
AssociateIspWithSkillV1 Associates an in-skill product with a skill.

	productId - The in-skill product ID.
	skillId - The skill ID.
*/
func (s *Client) AssociateIspWithSkillV1(productId string, skillId string) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/inSkillProducts/{productId}/skills/{skillId}")
	h.Path("productId", productId)
	h.Path("skillId", skillId)
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
 "description": "Associates an in-skill product with a skill.",
 "parameters": [
  {
   "description": "The in-skill product ID.",
   "in": "path",
   "name": "productId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "Success. No content.",
   "headers": {}
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
 "x-operation-name": "associateIspWithSkillV1"
}
*/
