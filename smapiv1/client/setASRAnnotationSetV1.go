package client

import (
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetASRAnnotationSetV1 API which updates the ASR annotation set properties. Currently, the only data can be updated is annotation set name.

	skillId - The skill ID.
	annotationSetId - Identifier of the ASR annotation set.
	updateAsrAnnotationSetPropertiesRequestV1 - Payload sent to the update ASR annotation set properties API.
*/
func (s *Client) SetASRAnnotationSetV1(skillId string, annotationSetId string, updateAsrAnnotationSetPropertiesRequestV1 *annotationSets_.UpdateAsrAnnotationSetPropertiesRequestObject) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets/{annotationSetId}")
	h.Path("skillId", skillId)
	h.Path("annotationSetId", annotationSetId)
	h.Body = updateAsrAnnotationSetPropertiesRequestV1
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which updates the ASR annotation set properties. Currently, the only data can be updated is annotation set name.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the ASR annotation set.",
   "in": "path",
   "name": "annotationSetId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Payload sent to the update ASR annotation set properties API.",
   "in": "body",
   "name": "UpdateAsrAnnotationSetPropertiesRequestV1",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.UpdateAsrAnnotationSetPropertiesRequestObject"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "ASR annotation set exists and properties are updated successfully."
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
 "summary": "update the ASR annotation set properties.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "setASRAnnotationSetV1"
}
*/
