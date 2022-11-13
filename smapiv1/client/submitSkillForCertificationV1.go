package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SubmitSkillForCertificationV1 Submit the skill for certification.

	skillId - The skill ID.
	submitSkillForCertificationRequest - Defines the request body for submitSkillForCertification API.
*/
func (s *Client) SubmitSkillForCertificationV1(skillId string, submitSkillForCertificationRequest *skill_.SubmitSkillForCertificationRequest) (err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/submit")
	h.Path("skillId", skillId)
	h.Body = submitSkillForCertificationRequest
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &skill_.StandardizedError{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &skill_.StandardizedError{})
	h.ResponseType(429, &skill_.StandardizedError{})
	h.ResponseType(500, &skill_.StandardizedError{})
	h.ResponseType(503, &skill_.StandardizedError{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Submit the skill for certification.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the request body for submitSkillForCertification API.",
   "in": "body",
   "name": "submitSkillForCertificationRequest",
   "required": false,
   "schema": {
    "$ref": "#/definitions/v1.skill.SubmitSkillForCertificationRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Success. There is no content but returns Location in the header.",
   "headers": {
    "Location": {
     "description": "Certification resource URL.",
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
 "x-operation-name": "submitSkillForCertificationV1"
}
*/
