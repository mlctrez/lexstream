package alexahosted

/*
HostedSkillMetadata Alexa Hosted skill's metadata
*/
type HostedSkillMetadata struct {
	AlexaHosted *HostedSkillInfo `json,omitempty:"alexaHosted"`
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
