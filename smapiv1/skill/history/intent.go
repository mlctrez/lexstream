package history

/*
Intent Provides the intent name, slots and confidence of the intent used in this interaction.
*/
type Intent struct {
	// The hypothesized intent for this utterance.
	Name       string      `json:"name"`
	Confidence *Confidence `json:"confidence"`
	// The hypothesized slot(s) for this interaction.
	Slots map[string]Slot `json:"slots"`
}

/*
{
 "description": "Provides the intent name, slots and confidence of the intent used in this interaction.",
 "properties": {
  "confidence": {
   "$ref": "#/definitions/v1.skill.history.Confidence"
  },
  "name": {
   "description": "The hypothesized intent for this utterance.",
   "type": "string"
  },
  "slots": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.history.Slot"
   },
   "description": "The hypothesized slot(s) for this interaction.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
