package interactionmodel

/*
ModelConfiguration Global configurations applicable to a skill's model.
*/
type ModelConfiguration struct {
	FallbackIntentSensitivity *FallbackIntentSensitivity `json:"fallbackIntentSensitivity"`
}

/*
{
 "description": "Global configurations applicable to a skill's model.",
 "properties": {
  "fallbackIntentSensitivity": {
   "$ref": "#/definitions/v1.skill.interactionModel.FallbackIntentSensitivity"
  }
 },
 "type": "object"
}
*/
