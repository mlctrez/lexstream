package alexahosted

/*
HostedSkillRepositoryInfo Alexa Hosted Skill's Repository Information
*/
type HostedSkillRepositoryInfo struct {
	Url   string                 `json:"url"`
	Type_ *HostedSkillRepository `json:"type"`
}

/*
{
 "description": "Alexa Hosted Skill's Repository Information",
 "properties": {
  "type": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepository",
   "x-isEnum": true
  },
  "url": {
   "type": "string"
  }
 },
 "required": [
  "type",
  "url"
 ],
 "type": "object"
}
*/
