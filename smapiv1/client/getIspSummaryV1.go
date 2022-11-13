package client

import (
	isp_ "github.com/mlctrez/lexstream/smapiv1/isp"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetIspSummaryV1 Get the summary information for an in-skill product.

	productId - The in-skill product ID.
	stage - Stage for skill.
*/
func (s *Client) GetIspSummaryV1(productId string, stage string) (response *isp_.InSkillProductSummaryResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/inSkillProducts/{productId}/stages/{stage}/summary")
	h.Path("productId", productId)
	h.Path("stage", stage)
	response = &isp_.InSkillProductSummaryResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the summary information for an in-skill product.",
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
   "description": "Returns current in-skill product summary for productId.",
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
    "$ref": "#/definitions/v1.isp.InSkillProductSummaryResponse"
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
 "x-operation-name": "getIspSummaryV1"
}
*/
