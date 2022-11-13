package client

import (
	annotationSets_ "github.com/mlctrez/lexstream/smapiv1/skill/asr/annotationSets"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetAnnotationsForASRAnnotationSetV1 API that updates the annotaions in the annotation set

	skillId - The skill ID.
	annotationSetId - Identifier of the ASR annotation set.
	updateAsrAnnotationSetContentsRequest - Payload containing annotation set contents. Two formats are accepted here: - `application/json`: Annotation set payload in JSON format. - `text/csv`: Annotation set payload in CSV format. Note that for CSV format, the first row should describe the column attributes. Columns should be delimited by comma.  The subsequent rows should describe annotation data and each annotation attributes has to follow the strict ordering defined in the first row. Each annotation fields should be delimited by comma.
*/
func (s *Client) SetAnnotationsForASRAnnotationSetV1(skillId string, annotationSetId string, updateAsrAnnotationSetContentsRequest *annotationSets_.UpdateAsrAnnotationSetContentsPayload) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/asrAnnotationSets/{annotationSetId}/annotations")
	h.Path("skillId", skillId)
	h.Path("annotationSetId", annotationSetId)
	h.Body = updateAsrAnnotationSetContentsRequest
	err = h.Execute(s.Client)
	return
}

/*
{
 "consumes": [
  "application/json",
  "text/csv"
 ],
 "description": "API that updates the annotaions in the annotation set\n",
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
   "description": "Payload containing annotation set contents. Two formats are accepted here: - `application/json`: Annotation set payload in JSON format. - `text/csv`: Annotation set payload in CSV format. Note that for CSV format, the first row should describe the column attributes. Columns should be delimited by comma.  The subsequent rows should describe annotation data and each annotation attributes has to follow the strict ordering defined in the first row. Each annotation fields should be delimited by comma.\n",
   "in": "body",
   "name": "UpdateAsrAnnotationSetContentsRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.UpdateAsrAnnotationSetContentsPayload"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "ASR annotation set contents have been updated successfully."
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
 "summary": "Update the annotations in the annotation set",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "setAnnotationsForASRAnnotationSetV1"
}
*/
