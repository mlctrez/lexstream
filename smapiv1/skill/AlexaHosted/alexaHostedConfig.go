package alexahosted

/*
AlexaHostedConfig Alexa hosted skill create configuration
*/
type AlexaHostedConfig struct {
	Runtime *HostedSkillRuntime `json,omitempty:"runtime"`
}

/*
{
 "description": "Alexa hosted skill create configuration",
 "properties": {
  "runtime": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRuntime",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
