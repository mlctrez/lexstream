package client

import (
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetAnnotationsForASRAnnotationSetV1

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a paginationContext.

	annotationSetId - Identifier of the ASR annotation set.
	accept - - `application/json`: indicate to download annotation set contents in JSON format - `text/csv`: indicate to download annotation set contents in CSV format
*/
func (s *Client) GetAnnotationsForASRAnnotationSetV1(skillId string, nextToken string, maxResults int, annotationSetId string, accept string) (response *annotationSets_.GetAsrAnnotationSetAnnotationsResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets/{annotationSetId}/annotations")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	h.Path("annotationSetId", annotationSetId)
	h.Header("Accept", accept)
	response = &annotationSets_.GetAsrAnnotationSetAnnotationsResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
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
   "default": 1000,
   "description": "Sets the maximum number of results returned in the response body. Defaults to 1000. If more results are present, the response will contain a paginationContext.\n",
   "in": "query",
   "maximum": 1000,
   "minimum": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "description": "Identifier of the ASR annotation set.",
   "in": "path",
   "name": "annotationSetId",
   "required": true,
   "type": "string"
  },
  {
   "description": "- `application/json`: indicate to download annotation set contents in JSON format - `text/csv`: indicate to download annotation set contents in CSV format\n",
   "enum": [
    "application/json",
    "text/csv"
   ],
   "in": "header",
   "name": "Accept",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "The annotation set contents payload in specified format.  This API also supports pagination for annotation set contents requested in  `application/json` content type. Paginaiton for requested content  type `text/csv` is not supported. In this case, the nextToken and  maxResults query parameters would be ignored even if they are  specified as query parameters.\n",
   "headers": {
    "Content-Type": {
     "description": "- `text/csv`: indicate the annotation set contents in CSV format. The first \n              row of the csv contents defines the attribute names of annotation\n              properties. Attribute names are delimited by comma. The subsequent rows\n              define the annotations. Each row represents an annotation content. The\n              annotation properties in each row follow the strict ordering of the attribute\n              property names defined in the first row.\n- `application/json`: indicate the annotation set contents in JSON format\n",
     "enum": [
      "text/csv",
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.GetAsrAnnotationSetAnnotationsResponse"
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Define the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
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
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "default": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Download the annotation set contents.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getAnnotationsForASRAnnotationSetV1"
}
*/
