package history

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type IntentRequest struct {
	DialogAct         *DialogAct            `json:"dialogAct"`
	Intent            *Intent               `json:"intent"`
	InteractionType   *InteractionType      `json:"interactionType"`
	Locale            *IntentRequestLocales `json:"locale"`
	PublicationStatus *PublicationStatus    `json:"publicationStatus"`
	Stage             *smapiv1.StageType    `json:"stage"`
	// The transcribed user speech.
	UtteranceText string `json:"utteranceText"`
}

/*
{
 "properties": {
  "dialogAct": {
   "$ref": "#/definitions/v1.skill.history.DialogAct"
  },
  "intent": {
   "$ref": "#/definitions/v1.skill.history.Intent"
  },
  "interactionType": {
   "$ref": "#/definitions/v1.skill.history.InteractionType",
   "x-isEnum": true
  },
  "locale": {
   "$ref": "#/definitions/v1.skill.history.IntentRequestLocales",
   "x-isEnum": true
  },
  "publicationStatus": {
   "$ref": "#/definitions/v1.skill.history.PublicationStatus",
   "x-isEnum": true
  },
  "stage": {
   "$ref": "#/definitions/v1.StageType",
   "x-isEnum": true
  },
  "utteranceText": {
   "description": "The transcribed user speech.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
