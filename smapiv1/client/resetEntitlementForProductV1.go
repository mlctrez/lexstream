package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
ResetEntitlementForProductV1 Resets the entitlement(s) of the Product for the current user.

	productId - The in-skill product ID.
	stage - Stage for skill.
*/
func (s *Client) ResetEntitlementForProductV1(productId string, stage string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/inSkillProducts/{productId}/stages/{stage}/entitlement")
	h.Path("productId", productId)
	h.Path("stage", stage)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Resets the entitlement(s) of the Product for the current user.",
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
 "x-operation-name": "resetEntitlementForProductV1"
}
*/
