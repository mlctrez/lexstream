package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/nlu/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateNLUAnnotationSetV1 This is an API that creates a new NLU annotation set with properties and returns the annotationId.

	skillId - The skill ID.
	createNLUAnnotationSetRequest - Payload sent to the create NLU annotation set API.
*/
func (s *Client) CreateNLUAnnotationSetV1(skillId string, createNLUAnnotationSetRequest *annotationSets_.CreateNLUAnnotationSetRequest) (response *annotationSets_.CreateNLUAnnotationSetResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets")
	h.Path("skillId", skillId)
	h.Body = createNLUAnnotationSetRequest
	response = &annotationSets_.CreateNLUAnnotationSetResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(500, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an API that creates a new NLU annotation set with properties and returns the annotationId.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the create NLU annotation set API.",
   "in": "body",
   "name": "CreateNLUAnnotationSetRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.CreateNLUAnnotationSetRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "NLU annotation set created successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    },
    "Location": {
     "description": "GET Location of the created NLU annotation set.\n",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.nlu.annotationSets.CreateNLUAnnotationSetResponse"
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
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Create a new NLU annotation set for a skill which will generate a new annotationId.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createNLUAnnotationSetV1"
}
*/
