package alexahosted

type HostedSkillInfo struct {
	Repository *HostedSkillRepositoryInfo `json,omitempty:"repository"`
	Runtime    *HostedSkillRuntime        `json,omitempty:"runtime"`
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
