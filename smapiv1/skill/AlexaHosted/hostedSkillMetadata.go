package alexahosted

/*
HostedSkillMetadata Alexa Hosted skill's metadata
*/
type HostedSkillMetadata struct {
	AlexaHosted *HostedSkillInfo `json:"alexaHosted"`
}

/*
{
 "description": "Alexa Hosted skill's metadata",
 "properties": {
  "alexaHosted": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillInfo"
  }
 },
 "type": "object"
}
*/
