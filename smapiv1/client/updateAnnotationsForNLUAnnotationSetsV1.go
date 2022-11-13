package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
UpdateAnnotationsForNLUAnnotationSetsV1 API which replaces the annotations in NLU annotation set.

	skillId - The skill ID.
	annotationId - Identifier of the NLU annotation set.
	content_Type - Standard HTTP. Pass `application/json` or `test/csv` for POST calls with a json/csv body.

	updateNLUAnnotationSetAnnotationsRequest - Payload sent to the update NLU annotation set API.
*/
func (s *Client) UpdateAnnotationsForNLUAnnotationSetsV1(skillId string, annotationId string, content_Type string, updateNLUAnnotationSetAnnotationsRequest *annotationSets_.UpdateNLUAnnotationSetAnnotationsRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets/{annotationId}/annotations")
	h.Path("skillId", skillId)
	h.Path("annotationId", annotationId)
	h.Header("Content-Type", content_Type)
	h.Body = updateNLUAnnotationSetAnnotationsRequest
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "consumes": [
  "application/json",
  "text/csv"
 ],
 "description": "API which replaces the annotations in NLU annotation set.\n",
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
   "description": "Standard HTTP. Pass `application/json` or `test/csv` for POST calls with a json/csv body.\n",
   "in": "header",
   "name": "Content-Type",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the update NLU annotation set API.",
   "in": "body",
   "name": "UpdateNLUAnnotationSetAnnotationsRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.UpdateNLUAnnotationSetAnnotationsRequest"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "NLU annotation set exists and starts the update.",
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
 "summary": "Replace the annotations in NLU annotation set.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "updateAnnotationsForNLUAnnotationSetsV1"
}
*/
