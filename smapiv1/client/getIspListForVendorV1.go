package client

import (
	isp_ "github.com/mlctrez/lexstream/smapiv1/isp"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetIspListForVendorV1 Get the list of in-skill products for the vendor.

	vendorId - The vendor ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	productId - The list of in-skill product IDs that you wish to get the summary for. A maximum of 50 in-skill product IDs can be specified in a single listInSkillProducts call. Please note that this parameter must not be used with 'nextToken' and/or 'maxResults' parameter.
	stage - Filter in-skill products by specified stage.
	type_ - Type of in-skill product to filter on.
	referenceName - Filter in-skill products by reference name.
	status - Status of in-skill product.
	isAssociatedWithSkill - Filter in-skill products by whether or not they are associated to a skill.
*/
func (s *Client) GetIspListForVendorV1(vendorId string, nextToken string, maxResults int, productId []string, stage string, type_ string, referenceName string, status string, isAssociatedWithSkill string) (response *isp_.ListInSkillProductResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/inSkillProducts")
	h.Param("vendorId", vendorId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	h.Param("productId", productId)
	h.Param("stage", stage)
	h.Param("type_", type_)
	h.Param("referenceName", referenceName)
	h.Param("status", status)
	h.Param("isAssociatedWithSkill", isAssociatedWithSkill)
	response = &isp_.ListInSkillProductResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get the list of in-skill products for the vendor.",
 "parameters": [
  {
   "description": "The vendor ID.",
   "in": "query",
   "name": "vendorId",
   "required": true,
   "type": "string"
  },
  {
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.",
   "in": "query",
   "maximum": 50,
   "minimum": 1,
   "multipleOf": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "collectionFormat": "multi",
   "description": "The list of in-skill product IDs that you wish to get the summary for. A maximum of 50 in-skill product IDs can be specified in a single listInSkillProducts call. Please note that this parameter must not be used with 'nextToken' and/or 'maxResults' parameter.",
   "in": "query",
   "items": {
    "type": "string"
   },
   "name": "productId",
   "required": false,
   "type": "array"
  },
  {
   "description": "Filter in-skill products by specified stage.",
   "enum": [
    "development",
    "live"
   ],
   "in": "query",
   "name": "stage",
   "required": false,
   "type": "string"
  },
  {
   "description": "Type of in-skill product to filter on.",
   "enum": [
    "SUBSCRIPTION",
    "ENTITLEMENT",
    "CONSUMABLE"
   ],
   "in": "query",
   "name": "type",
   "required": false,
   "type": "string"
  },
  {
   "description": "Filter in-skill products by reference name.",
   "in": "query",
   "name": "referenceName",
   "required": false,
   "type": "string"
  },
  {
   "description": "Status of in-skill product.",
   "enum": [
    "INCOMPLETE",
    "COMPLETE",
    "CERTIFICATION",
    "PUBLISHED",
    "SUPPRESSED"
   ],
   "in": "query",
   "name": "status",
   "required": false,
   "type": "string"
  },
  {
   "description": "Filter in-skill products by whether or not they are associated to a skill.",
   "enum": [
    "ASSOCIATED_WITH_SKILL",
    "NO_SKILL_ASSOCIATIONS"
   ],
   "in": "query",
   "name": "isAssociatedWithSkill",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Response contains list of in-skill products for the specified vendor and stage.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.isp.ListInSkillProductResponse"
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
 "x-operation-name": "getIspListForVendorV1"
}
*/
