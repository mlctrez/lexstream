package client

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	evaluations_ "github.com/mlctrez/lexstream/smapiv1/skill/evaluations"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
ProfileNluV1 This is a synchronous API that profiles an utterance against interaction model.

	profileNluRequest - Payload sent to the profile nlu API.
	skillId - The skill ID.
	stage - Stage for skill.
	locale - The locale for the model requested e.g. en-GB, en-US, de-DE.
*/
func (s *Client) ProfileNluV1(profileNluRequest *evaluations_.ProfileNluRequest, skillId string, stage string, locale string) (response *evaluations_.ProfileNluResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills/{skillId}/stages/{stage}/interactionModel/locales/{locale}/profileNlu")
	h.Body = profileNluRequest
	h.Path("skillId", skillId)
	h.Path("stage", stage)
	h.Path("locale", locale)
	response = &evaluations_.ProfileNluResponse{}
	h.Response = response
	h.ResponseType(400, &smapiv1.BadRequestError{})
	h.ResponseType(401, &smapiv1.Error{})
	h.ResponseType(403, &smapiv1.BadRequestError{})
	h.ResponseType(409, &smapiv1.Error{})
	h.ResponseType(429, &smapiv1.Error{})
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "This is a synchronous API that profiles an utterance against interaction model.",
 "parameters": [
  {
   "description": "Payload sent to the profile nlu API.",
   "in": "body",
   "name": "profileNluRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.evaluations.ProfileNluRequest"
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
  },
  {
   "description": "The locale for the model requested e.g. en-GB, en-US, de-DE.",
   "format": "languager-region; same as BCP-47 language tag format",
   "in": "path",
   "name": "locale",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Profiled utterance against interaction model and returned nlu response successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.evaluations.ProfileNluResponse"
   }
  },
  "400": {
   "description": "Bad request due to invalid or missing data.",
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
   "description": "The operation being requested is not allowed.",
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
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   }
  }
 },
 "summary": "Profile a test utterance.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "profileNluV1"
}
*/
