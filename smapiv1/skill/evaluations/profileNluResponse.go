package evaluations

type ProfileNluResponse struct {
	/*
	   Represents when an utterance results in the skill exiting. It would be true when NLU selects 1P ExitAppIntent or GoHomeIntent, and false otherwise.
	*/
	SessionEnded   bool                      `json:"sessionEnded,omitempty"`
	SelectedIntent *ProfileNluSelectedIntent `json:"selectedIntent,omitempty"`
	// All intents that Alexa considered for the utterance in the request, but did not select.
	ConsideredIntents []*Intent  `json:"consideredIntents,omitempty"`
	MultiTurn         *MultiTurn `json:"multiTurn,omitempty"`
}

/*
{
 "properties": {
  "consideredIntents": {
   "description": "All intents that Alexa considered for the utterance in the request, but did not select.",
   "items": {
    "$ref": "#/definitions/v1.skill.evaluations.Intent"
   },
   "type": "array"
  },
  "multiTurn": {
   "$ref": "#/definitions/v1.skill.evaluations.MultiTurn"
  },
  "selectedIntent": {
   "$ref": "#/definitions/v1.skill.evaluations.ProfileNluSelectedIntent"
  },
  "sessionEnded": {
   "description": "Represents when an utterance results in the skill exiting. It would be true when NLU selects 1P ExitAppIntent or GoHomeIntent, and false otherwise.\n",
   "type": "boolean"
  }
 },
 "type": "object"
}
*/
