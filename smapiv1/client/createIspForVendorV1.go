package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	isp_ "github.com/mlctrez/lexstream/smapiv1/isp"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateIspForVendorV1 Creates a new in-skill product for given vendorId.

	createInSkillProductRequest - defines the request body for createInSkillProduct API.
*/
func (s *Client) CreateIspForVendorV1(createInSkillProductRequest *isp_.CreateInSkillProductRequest) (response *isp_.ProductResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/inSkillProducts")
	h.Body = createInSkillProductRequest
	response = &isp_.ProductResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new in-skill product for given vendorId.",
 "parameters": [
  {
   "description": "defines the request body for createInSkillProduct API.",
   "in": "body",
   "name": "createInSkillProductRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.isp.createInSkillProductRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Success.",
   "headers": {},
   "schema": {
    "$ref": "#/definitions/v1.isp.ProductResponse"
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
 "x-operation-name": "createIspForVendorV1"
}
*/
