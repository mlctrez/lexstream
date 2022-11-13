package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateASRAnnotationSetV1 This is an API that creates a new ASR annotation set with a name and returns the annotationSetId which can later be used to retrieve or reference the annotation set

	skillId - The skill ID.
	createAsrAnnotationSetRequest - Payload sent to the create ASR annotation set API.
*/
func (s *Client) CreateASRAnnotationSetV1(skillId string, createAsrAnnotationSetRequest *annotationSets_.CreateAsrAnnotationSetRequestObject) (response *annotationSets_.CreateAsrAnnotationSetResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets")
	h.Path("skillId", skillId)
	h.Body = createAsrAnnotationSetRequest
	response = &annotationSets_.CreateAsrAnnotationSetResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	h.ResponseType(503, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an API that creates a new ASR annotation set with a name and returns the annotationSetId which can later be used to retrieve or reference the annotation set\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the create ASR annotation set API.",
   "in": "body",
   "name": "CreateAsrAnnotationSetRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.CreateAsrAnnotationSetRequestObject"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "ASR annotation set created successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    },
    "Location": {
     "description": "Location of the created ASR annotation set.\n",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.CreateAsrAnnotationSetResponse"
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
 "summary": "Create a new ASR annotation set for a skill",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "createASRAnnotationSetV1"
}
*/
