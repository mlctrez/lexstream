package interactionmodel

/*
FallbackIntentSensitivity Denotes skill's sensitivity for out-of-domain utterances.
*/
type FallbackIntentSensitivity struct {
	Level *FallbackIntentSensitivityLevel `json:"level,omitempty"`
}

/*
{
 "description": "Denotes skill's sensitivity for out-of-domain utterances.",
 "properties": {
  "level": {
   "$ref": "#/definitions/v1.skill.interactionModel.FallbackIntentSensitivityLevel",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
