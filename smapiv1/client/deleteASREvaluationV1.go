package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
DeleteASREvaluationV1 API which enables the deletion of an evaluation.

	skillId - The skill ID.
	evaluationId - Identifier of the evaluation.
*/
func (s *Client) DeleteASREvaluationV1(skillId string, evaluationId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/{skillId}/asrEvaluations/{evaluationId}")
	h.Path("skillId", skillId)
	h.Path("evaluationId", evaluationId)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "API which enables the deletion of an evaluation. \n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the evaluation.",
   "in": "path",
   "name": "evaluationId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "ASR evaluation exists and is deleted successfully."
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
 "summary": "Delete an evaluation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "deleteASREvaluationV1"
}
*/
