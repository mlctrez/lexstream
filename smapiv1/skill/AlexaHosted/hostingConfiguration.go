package alexahosted

/*
HostingConfiguration Configurations for creating new hosted skill
*/
type HostingConfiguration struct {
	AlexaHosted *AlexaHostedConfig `json:"alexaHosted,omitempty"`
}

/*
{
 "description": "Configurations for creating new hosted skill",
 "properties": {
  "alexaHosted": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.AlexaHostedConfig"
  }
 },
 "type": "object"
}
*/
