package manifest

/*
LocalizedHealthInfo Defines the structure for health skill locale specific publishing information in the skill manifest.
*/
type LocalizedHealthInfo struct {
	// SSML supported name to use when Alexa renders the health skill name in a prompt to the user.
	PromptName string `json,omitempty:"promptName"`
	// Defines the names to use when a user tries to invoke the health skill.
	Aliases []*HealthAlias `json,omitempty:"aliases"`
}

/*
{
 "description": "Defines the structure for health skill locale specific publishing information in the skill manifest.",
 "properties": {
  "aliases": {
   "description": "Defines the names to use when a user tries to invoke the health skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.HealthAlias"
   },
   "type": "array"
  },
  "promptName": {
   "description": "SSML supported name to use when Alexa renders the health skill name in a prompt to the user.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
