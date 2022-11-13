package client

import (
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdatePropertiesForNLUAnnotationSetsV1 API which updates the NLU annotation set properties. Currently, the only data can be updated is annotation set name.

	skillId - The skill ID.
	annotationId - Identifier of the NLU annotation set.
	updateNLUAnnotationSetPropertiesRequest - Payload sent to the update NLU annotation set properties API.
*/
func (s *Client) UpdatePropertiesForNLUAnnotationSetsV1(skillId string, annotationId string, updateNLUAnnotationSetPropertiesRequest *annotationSets_.UpdateNLUAnnotationSetPropertiesRequest) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets/{annotationId}/properties")
	h.Path("skillId", skillId)
	h.Path("annotationId", annotationId)
	h.Body = updateNLUAnnotationSetPropertiesRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which updates the NLU annotation set properties. Currently, the only data can be updated is annotation set name.\n",
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
  },
  {
   "description": "Payload sent to the update NLU annotation set properties API.",
   "in": "body",
   "name": "UpdateNLUAnnotationSetPropertiesRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.UpdateNLUAnnotationSetPropertiesRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "NLU annotation set exists and properties are updated successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/hal+json or application/json supported",
     "type": "string"
    }
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
 "summary": "update the NLU annotation set properties.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "updatePropertiesForNLUAnnotationSetsV1"
}
*/
