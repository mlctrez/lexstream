package alexahosted

type HostedSkillInfo struct {
	Repository *HostedSkillRepositoryInfo `json:"repository,omitempty"`
	Runtime    *HostedSkillRuntime        `json:"runtime,omitempty"`
}

/*
{
 "properties": {
  "repository": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepositoryInfo"
  },
  "runtime": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRuntime",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
