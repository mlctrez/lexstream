package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	vendorManagement_ "github.com/mlctrez/lexstream/smapiv1/vendorManagement"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetVendorListV1 Get the list of Vendor information.
*/
func (s *Client) GetVendorListV1() (response *vendorManagement_.Vendors, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/vendors")
	response = &vendorManagement_.Vendors{}
	h.Response = response
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the list of Vendor information.\n",
 "parameters": [],
 "responses": {
  "200": {
   "description": "Return vendor information on success.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.vendorManagement.Vendors"
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
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getVendorListV1"
}
*/
