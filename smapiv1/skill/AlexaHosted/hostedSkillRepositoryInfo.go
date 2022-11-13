package alexahosted

/*
HostedSkillRepositoryInfo Alexa Hosted Skill's Repository Information
*/
type HostedSkillRepositoryInfo struct {
	Url   string                 `json:"url,omitempty"`
	Type_ *HostedSkillRepository `json:"type,omitempty"`
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
