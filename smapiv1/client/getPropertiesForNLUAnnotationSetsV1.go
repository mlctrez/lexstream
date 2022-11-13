package client

import (
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetPropertiesForNLUAnnotationSetsV1 Return the properties for an NLU annotation set.

	skillId - The skill ID.
	annotationId - Identifier of the NLU annotation set.
*/
func (s *Client) GetPropertiesForNLUAnnotationSetsV1(skillId string, annotationId string) (response *annotationSets_.GetNLUAnnotationSetPropertiesResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets/{annotationId}/properties")
	h.Path("skillId", skillId)
	h.Path("annotationId", annotationId)
	response = &annotationSets_.GetNLUAnnotationSetPropertiesResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Return the properties for an NLU annotation set.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the NLU annotation set.",
   "in": "path",
   "name": "annotationId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "The NLU annotation set exists.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/hal+json or application/json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.GetNLUAnnotationSetPropertiesResponse"
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
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Get the properties of an NLU annotation set",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getPropertiesForNLUAnnotationSetsV1"
}
*/
