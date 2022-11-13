package evaluations

type ProfileNluRequest struct {
	// Actual representation of user input to Alexa.
	Utterance string `json:"utterance"`
	// Opaque string which contains multi-turn related context.
	MultiTurnToken string `json:"multiTurnToken"`
}

/*
{
 "properties": {
  "multiTurnToken": {
   "description": "Opaque string which contains multi-turn related context.",
   "type": "string"
  },
  "utterance": {
   "description": "Actual representation of user input to Alexa.",
   "type": "string"
  }
 },
 "required": [
  "utterance"
 ],
 "type": "object"
}
*/
