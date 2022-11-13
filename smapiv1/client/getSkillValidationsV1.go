package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	validations_ "github.com/mlctrez/lexstream/smapiv1/skill/validations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSkillValidationsV1 This API gets the result of a previously executed validation. A successful response will contain the status of the executed validation. If the validation successfully completed, the response will also contain information related to executed validations. In cases where requests to this API results in an error, the response will contain a description of the problem. In cases where the validation failed, the response will contain a status attribute indicating that a failure occurred. Note that validation results are stored for 60 minutes. A request for an expired validation result will return a 404 HTTP status code.

	skillId - The skill ID.
	validationId - Id of the validation. Reserved word identifier of mostRecent can be used to get the most recent validation for the skill and stage. Note that the behavior of the API in this case would be the same as when the actual validation id of the most recent validation is used in the request.

	stage - Stage for skill.
	accept_Language - User's locale/language in context.
*/
func (s *Client) GetSkillValidationsV1(skillId string, validationId string, stage string, accept_Language string) (response *validations_.ValidationsApiResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/validations/{validationId}")
	h.Path("skillId", skillId)
	h.Path("validationId", validationId)
	h.Path("stage", stage)
	h.Header("Accept-Language", accept_Language)
	response = &validations_.ValidationsApiResponse{}
	h.Response = response
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(404, &smapiv1.Error{})
	h.ResponseType(409, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This API gets the result of a previously executed validation. A successful response will contain the status of the executed validation. If the validation successfully completed, the response will also contain information related to executed validations. In cases where requests to this API results in an error, the response will contain a description of the problem. In cases where the validation failed, the response will contain a status attribute indicating that a failure occurred. Note that validation results are stored for 60 minutes. A request for an expired validation result will return a 404 HTTP status code.\n",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Id of the validation. Reserved word identifier of mostRecent can be used to get the most recent validation for the skill and stage. Note that the behavior of the API in this case would be the same as when the actual validation id of the most recent validation is used in the request.\n",
   "in": "path",
   "name": "validationId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Stage for skill.",
   "in": "path",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "User's locale/language in context.",
   "in": "header",
   "name": "Accept-Language",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successfully retrieved skill validation information.",
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
    "$ref": "#/definitions/v1.skill.validations.ValidationsApiResponse"
   }
  },
  "403": {
   "description": "API user does not have permission or is currently in a state that does not allow calls to this API.\n",
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
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The specified skill, stage, or validation does not exist. The error response will contain a description that indicates the specific resource type that was not found.\n",
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
   }
  }
 },
 "summary": "Get the result of a previously executed validation.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getSkillValidationsV1"
}
*/
