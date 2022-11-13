package skill

/*
SkillCredentials Structure for skill credentials response.
*/
type SkillCredentials struct {
	SkillMessagingCredentials *SkillMessagingCredentials `json:"skillMessagingCredentials"`
}

/*
{
 "description": "Structure for skill credentials response.",
 "properties": {
  "skillMessagingCredentials": {
   "$ref": "#/definitions/v1.skill.SkillMessagingCredentials"
  }
 },
 "type": "object"
}
*/
