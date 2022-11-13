package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
DeleteIspForProductV1 Deletes the in-skill product for given productId. Only development stage supported. Live in-skill products or in-skill products associated with a skill cannot be deleted by this API.

	productId - The in-skill product ID.
	stage - Stage for skill.
	if_Match - Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.
*/
func (s *Client) DeleteIspForProductV1(productId string, stage string, if_Match string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/inSkillProducts/{productId}/stages/{stage}")
	h.Path("productId", productId)
	h.Path("stage", stage)
	h.Header("If-Match", if_Match)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Deletes the in-skill product for given productId. Only development stage supported. Live in-skill products or in-skill products associated with a skill cannot be deleted by this API.",
 "parameters": [
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
   "description": "Request header that specified an entity tag. The server will update the resource only if the eTag matches with the resource's current eTag.",
   "in": "header",
   "name": "If-Match",
   "required": false,
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
 "x-operation-name": "deleteIspForProductV1"
}
*/
