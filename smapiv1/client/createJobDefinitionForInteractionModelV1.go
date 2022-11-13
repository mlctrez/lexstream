package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill "github.com/mlctrez/lexstream/smapiv1/skill"
	jobs_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/jobs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
CreateJobDefinitionForInteractionModelV1 Creates a new Job Definition from the Job Definition request provided. This can be either a CatalogAutoRefresh, which supports time-based configurations for catalogs, or a ReferencedResourceVersionUpdate, which is used for slotTypes and Interaction models to be automatically updated on the dynamic update of their referenced catalog.

	createJobDefinitionRequest - Request to create a new Job Definition.
*/
func (s *Client) CreateJobDefinitionForInteractionModelV1(createJobDefinitionRequest *jobs_.CreateJobDefinitionRequest) (response *jobs_.CreateJobDefinitionResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/api/custom/interactionModel/jobs")
	h.Body = createJobDefinitionRequest
	response = &jobs_.CreateJobDefinitionResponse{}
	h.Response = response
	h.ResponseType(400, &jobs_.ValidationErrors{})
	h.ResponseType(401, &skill.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(429, &skill.StandardizedError{})
	h.ResponseType(500, &skill.StandardizedError{})
	h.ResponseType(503, &skill.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Creates a new Job Definition from the Job Definition request provided. This can be either a CatalogAutoRefresh, which supports time-based configurations for catalogs, or a ReferencedResourceVersionUpdate, which is used for slotTypes and Interaction models to be automatically updated on the dynamic update of their referenced catalog.\n",
 "parameters": [
  {
   "description": "Request to create a new Job Definition.",
   "in": "body",
   "name": "createJobDefinitionRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.CreateJobDefinitionRequest"
   }
  }
 ],
 "responses": {
  "201": {
   "description": "Returns the generated jobId.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/hal+json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.CreateJobDefinitionResponse"
   }
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
 "x-operation-name": "createJobDefinitionForInteractionModelV1"
}
*/
