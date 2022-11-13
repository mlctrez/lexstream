package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	version_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/version"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ListInteractionModelCatalogVersionsV1 List all the historical versions of the given catalogId.

	catalogId - Provides a unique identifier of the catalog
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	sortDirection - Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.
	sortField - Sets the field on which the sorting would be applied.
*/
func (s *Client) ListInteractionModelCatalogVersionsV1(catalogId string, maxResults int, nextToken string, sortDirection string, sortField string) (response *version_.ListCatalogEntityVersionsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/api/custom/interactionModel/catalogs/{catalogId}/versions")
	h.Path("catalogId", catalogId)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	h.Param("sortDirection", sortDirection)
	h.Param("sortField", sortField)
	response = &version_.ListCatalogEntityVersionsResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "List all the historical versions of the given catalogId.",
 "parameters": [
  {
   "description": "Provides a unique identifier of the catalog",
   "in": "path",
   "name": "catalogId",
   "required": true,
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
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.",
   "in": "query",
   "name": "sortDirection",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the field on which the sorting would be applied.",
   "in": "query",
   "name": "sortField",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns list of catalogs for the vendor.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.ListCatalogEntityVersionsResponse"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. the catalog definition is invalid.",
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified catalog does not exist.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "listInteractionModelCatalogVersionsV1"
}
*/
