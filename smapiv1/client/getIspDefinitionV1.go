package client

import (
	isp_ "github.com/mlctrez/lexstream/smapiv1/isp"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetIspDefinitionV1 Returns the in-skill product definition for given productId.

	productId - The in-skill product ID.
	stage - Stage for skill.
*/
func (s *Client) GetIspDefinitionV1(productId string, stage string) (response *isp_.InSkillProductDefinitionResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/inSkillProducts/{productId}/stages/{stage}")
	h.Path("productId", productId)
	h.Path("stage", stage)
	response = &isp_.InSkillProductDefinitionResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Returns the in-skill product definition for given productId.",
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
  "200": {
   "description": "Response contains the latest version of an in-skill product for the specified stage.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    },
    "ETag": {
     "description": "Identifier for the version of the resource, can be used for conditional updates.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.isp.InSkillProductDefinitionResponse"
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
 "x-operation-name": "getIspDefinitionV1"
}
*/
