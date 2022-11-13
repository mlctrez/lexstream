package client

import (
	validations_ "github.com/mlctrez/lexstream/smapiv1/skill/validations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SubmitSkillValidationV1 This is an asynchronous API which allows a skill developer to execute various validations against their skill.

	validationsApiRequest - Payload sent to the skill validation API.
	skillId - The skill ID.
	stage - Stage for skill.
*/
func (s *Client) SubmitSkillValidationV1(validationsApiRequest *validations_.ValidationsApiRequest, skillId string, stage string) (response *validations_.ValidationsApiResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/validations")
	h.Body = validationsApiRequest
	h.Path("skillId", skillId)
	h.Path("stage", stage)
	response = &validations_.ValidationsApiResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is an asynchronous API which allows a skill developer to execute various validations against their skill.\n",
 "parameters": [
  {
   "description": "Payload sent to the skill validation API.",
   "in": "body",
   "name": "validationsApiRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.validations.ValidationsApiRequest"
   }
  },
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage for skill.",
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "202": {
   "description": "Skill validation has successfully begun.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Path to validation resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.validations.ValidationsApiResponse"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "API user does not have permission or is currently in a state that does not allow calls to this API.\n",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill, stage or validation does not exist. The error response will contain a description that indicates the specific resource type that was not found.\n",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "409": {
   "description": "This requests conflicts with another one currently being processed.\n",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "API user has exceeded the permitted request rate.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal service error.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. It falls back to en-US if the locale in the request is not supported.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Validate a skill.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "submitSkillValidationV1"
}
*/
