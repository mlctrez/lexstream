package client

import (
	history_ "github.com/mlctrez/lexstream/smapiv1/skill/history"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetUtteranceDataV1

	skillId - The skill ID.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	sortDirection - Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.
	sortField - Sets the field on which the sorting would be applied.
	stage - The stage of the skill to be used for evaluation. An error will be returned if this skill stage is not enabled on the account used for evaluation.
	locale -
	dialogActname - A filter used to retrieve items where the dialogAct name is equal to the given value.

* `Dialog.ElicitSlot`: Alexa asked the user for the value of a specific slot. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#elicitslot)
* `Dialog.ConfirmSlot`: Alexa confirmed the value of a specific slot before continuing with the dialog. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmslot)
* `Dialog.ConfirmIntent`: Alexa confirmed the all the information the user has provided for the intent before the skill took action. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmintent)

	intentconfidencebin -
	intentname - A filter used to retrieve items where the intent name is equal to the given value.
	intentslotsname - A filter used to retrieve items where the one of the slot names is equal to the given value.
	interactionType -
	publicationStatus -
	utteranceText - A filter used to retrieve items where the utterance text contains the given phrase. Each filter value can be at-least 1 character and at-most 100 characters long.
*/
func (s *Client) GetUtteranceDataV1(skillId string, nextToken string, maxResults int, sortDirection string, sortField string, stage string, locale []history_.LocaleInQuery, dialogActname []history_.DialogActName, intentconfidencebin []history_.IntentConfidenceBin, intentname []string, intentslotsname []string, interactionType []history_.InteractionType, publicationStatus []history_.PublicationStatus, utteranceText []string) (response *history_.IntentRequests, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/history/intentRequests")
	h.Path("skillId", skillId)
	h.Param("nextToken", nextToken)
	h.Param("maxResults", maxResults)
	h.Param("sortDirection", sortDirection)
	h.Param("sortField", sortField)
	h.Param("stage", stage)
	h.Param("locale", locale)
	h.Param("dialogActname", dialogActname)
	h.Param("intentconfidencebin", intentconfidencebin)
	h.Param("intentname", intentname)
	h.Param("intentslotsname", intentslotsname)
	h.Param("interactionType", interactionType)
	h.Param("publicationStatus", publicationStatus)
	h.Param("utteranceText", utteranceText)
	response = &history_.IntentRequests{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.",
   "in": "query",
   "maximum": 50,
   "minimum": 1,
   "multipleOf": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "description": "Sets the sorting direction of the result items. When set to 'asc' these items are returned in ascending order of sortField value and when set to 'desc' these items are returned in descending order of sortField value.",
   "in": "query",
   "name": "sortDirection",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the field on which the sorting would be applied.",
   "in": "query",
   "name": "sortField",
   "required": false,
   "type": "string"
  },
  {
   "description": "The stage of the skill to be used for evaluation. An error will be returned if this skill stage is not enabled on the account used for evaluation.",
   "enum": [
    "development",
    "live"
   ],
   "in": "query",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "in": "query",
   "items": {
    "$ref": "#/definitions/v1.skill.history.localeInQuery"
   },
   "name": "locale",
   "required": false,
   "type": "array"
  },
  {
   "description": "A filter used to retrieve items where the dialogAct name is equal to the given value.\n* `Dialog.ElicitSlot`: Alexa asked the user for the value of a specific slot. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#elicitslot)\n* `Dialog.ConfirmSlot`: Alexa confirmed the value of a specific slot before continuing with the dialog. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmslot)\n* `Dialog.ConfirmIntent`: Alexa confirmed the all the information the user has provided for the intent before the skill took action. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmintent)\n",
   "in": "query",
   "items": {
    "$ref": "#/definitions/v1.skill.history.DialogActName"
   },
   "name": "dialogAct.name",
   "required": false,
   "type": "array"
  },
  {
   "in": "query",
   "items": {
    "$ref": "#/definitions/v1.skill.history.IntentConfidenceBin"
   },
   "name": "intent.confidence.bin",
   "required": false,
   "type": "array"
  },
  {
   "description": "A filter used to retrieve items where the intent name is equal to the given value.",
   "in": "query",
   "items": {
    "type": "string"
   },
   "name": "intent.name",
   "required": false,
   "type": "array"
  },
  {
   "description": "A filter used to retrieve items where the one of the slot names is equal to the given value.",
   "in": "query",
   "items": {
    "type": "string"
   },
   "name": "intent.slots.name",
   "required": false,
   "type": "array"
  },
  {
   "in": "query",
   "items": {
    "$ref": "#/definitions/v1.skill.history.InteractionType"
   },
   "name": "interactionType",
   "required": false,
   "type": "array"
  },
  {
   "in": "query",
   "items": {
    "$ref": "#/definitions/v1.skill.history.PublicationStatus"
   },
   "name": "publicationStatus",
   "required": false,
   "type": "array"
  },
  {
   "description": "A filter used to retrieve items where the utterance text contains the given phrase. Each filter value can be at-least 1 character and at-most 100 characters long.",
   "in": "query",
   "items": {
    "type": "string"
   },
   "name": "utteranceText",
   "required": false,
   "type": "array"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns a list of utterance items for the given skill.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json+hal supported",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.history.IntentRequests"
   }
  },
  "400": {
   "description": "Bad Request.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "Unauthorized.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "404": {
   "description": "Skill Not Found.",
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
 "summary": "The Intent Request History API provides customers with the aggregated and anonymized transcription of user speech data and intent request details for their skills.",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getUtteranceDataV1"
}
*/
