package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	jobs "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/jobs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CancelNextJobExecutionForInteractionModelV1 Cancel the next execution for the given job.

	jobId - The identifier for dynamic jobs.
	executionId - The identifier for dynamic job executions. Currently only allowed for scheduled executions.
*/
func (s *Client) CancelNextJobExecutionForInteractionModelV1(jobId string, executionId string) (err error) {
	h := swaggerlt.NewRequestHelper("delete", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs/{jobId}/executions/{executionId}")
	h.Path("jobId", jobId)
	h.Path("executionId", executionId)
	h.ResponseType(400, &jobs.ValidationErrors{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill.StandardizedError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Cancel the next execution for the given job.",
 "parameters": [
  {
   "description": "The identifier for dynamic jobs.",
   "in": "path",
   "name": "jobId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The identifier for dynamic job executions. Currently only allowed for scheduled executions.",
   "in": "path",
   "name": "executionId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "204": {
   "description": "No Content; Confirms that the next execution is canceled.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
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
 "x-operation-name": "cancelNextJobExecutionForInteractionModelV1"
}
*/
