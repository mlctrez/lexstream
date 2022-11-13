package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
DeleteJobDefinitionForInteractionModelV1 Delete the job definition for a given jobId.

	jobId - The identifier for dynamic jobs.
*/
func (s *Client) DeleteJobDefinitionForInteractionModelV1(jobId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs/{jobId}")
	h.Path("jobId", jobId)
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Delete the job definition for a given jobId.",
 "parameters": [
  {
   "description": "The identifier for dynamic jobs.",
   "in": "path",
   "name": "jobId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "No content, confirms the resource is updated."
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.ValidationErrors"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Defines the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
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
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "deleteJobDefinitionForInteractionModelV1"
}
*/
