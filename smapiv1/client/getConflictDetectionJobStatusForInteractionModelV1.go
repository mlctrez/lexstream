package client

import (
	conflictDetection_ "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel/conflictDetection"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetConflictDetectionJobStatusForInteractionModelV1 This API returns the job status of conflict detection job for a specified interaction model.

	skillId - The skill ID.
	locale - The locale for the model requested e.g. en-GB, en-US, de-DE.
	stage - Stage of the interaction model.
	version - Version of interaction model. Use "~current" to get the model of the current version.
*/
func (s *Client) GetConflictDetectionJobStatusForInteractionModelV1(skillId string, locale string, stage string, version string) (response *conflictDetection_.GetConflictDetectionJobStatusResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/interactionModel/locales/{locale}/versions/{version}/conflictDetectionJobStatus")
	h.Path("skillId", skillId)
	h.Path("locale", locale)
	h.Path("stage", stage)
	h.Path("version", version)
	response = &conflictDetection_.GetConflictDetectionJobStatusResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This API returns the job status of conflict detection job for a specified interaction model.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The locale for the model requested e.g. en-GB, en-US, de-DE.",
   "format": "languager-region; same as BCP-47 language tag format",
   "in": "path",
   "name": "locale",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage of the interaction model.",
   "enum": [
    "development"
   ],
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "Version of interaction model. Use \"~current\" to get the model of the current version.",
   "in": "path",
   "name": "version",
   "pattern": "^~current|[0-9]+$",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Get conflict detection results successfully.",
   "headers": {
    "Content-Type": {
     "description": "returned content type, only application/json supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.GetConflictDetectionJobStatusResponse"
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
   "description": "There is no catalog defined for the catalogId.",
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
 "summary": "Retrieve conflict detection job status for skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getConflictDetectionJobStatusForInteractionModelV1"
}
*/
