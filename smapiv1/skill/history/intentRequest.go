package history

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type IntentRequest struct {
	DialogAct         *DialogAct            `json,omitempty:"dialogAct"`
	Intent            *Intent               `json,omitempty:"intent"`
	InteractionType   *InteractionType      `json,omitempty:"interactionType"`
	Locale            *IntentRequestLocales `json,omitempty:"locale"`
	PublicationStatus *PublicationStatus    `json,omitempty:"publicationStatus"`
	Stage             *smapiv1.StageType    `json,omitempty:"stage"`
	// The transcribed user speech.
	UtteranceText string `json,omitempty:"utteranceText"`
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
