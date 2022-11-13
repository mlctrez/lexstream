package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	jobs_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/jobs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetJobStatusForInteractionModelV1 Update the JobStatus to Enable or Disable a job.

	jobId - The identifier for dynamic jobs.
	updateJobStatusRequest - Request to update Job Definition status.
*/
func (s *Client) SetJobStatusForInteractionModelV1(jobId string, updateJobStatusRequest *jobs_.UpdateJobStatusRequest) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs/{jobId}/status")
	h.Path("jobId", jobId)
	h.Body = updateJobStatusRequest
	h.ResponseType(400, &jobs_.ValidationErrors{})
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
 "description": "Update the JobStatus to Enable or Disable a job.",
 "parameters": [
  {
   "description": "The identifier for dynamic jobs.",
   "in": "path",
   "name": "jobId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Request to update Job Definition status.",
   "in": "body",
   "name": "updateJobStatusRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.UpdateJobStatusRequest"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "No content; Confirms that the fields are updated.",
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
     "description": "Contains content type of the response; only application/json supported.",
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
 "x-operation-name": "setJobStatusForInteractionModelV1"
}
*/
