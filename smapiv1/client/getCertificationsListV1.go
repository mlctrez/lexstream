package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	certification_ "github.com/mlctrez/lexstream/smapiv1/skill/certification"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetCertificationsListV1 Get list of all certifications available for a skill, including information about past certifications and any ongoing certification. The default sort order is descending on skillSubmissionTimestamp for Certifications.

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
*/
func (s *Client) GetCertificationsListV1(skillId string, nextToken string, maxResults int) (response *certification_.ListCertificationsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/certifications")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	response = &certification_.ListCertificationsResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get list of all certifications available for a skill, including information about past certifications and any ongoing certification. The default sort order is descending on skillSubmissionTimestamp for Certifications.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
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
  }
 ],
 "responses": {
  "200": {
   "description": "Returns list of certifications for the skillId.",
   "headers": {
    "Content-Type": {
     "description": "The content type of the response body. only application/hal+json supported.",
     "enum": [
      "application/hal+json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.certification.ListCertificationsResponse"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error e.g. if any request parameter is invalid like certification Id or pagination token etc. If the maxResults is not in the range of 1 to 50, it also qualifies for this error.\n",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceeded the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.\n",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getCertificationsListV1"
}
*/
