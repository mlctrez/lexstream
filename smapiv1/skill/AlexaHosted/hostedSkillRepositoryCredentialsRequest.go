package alexahosted

type HostedSkillRepositoryCredentialsRequest struct {
	Repository *HostedSkillRepositoryInfo `json:"repository"`
}

/*
{
 "properties": {
  "repository": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepositoryInfo"
  }
 },
 "required": [
  "repository"
 ],
 "type": "object"
}
*/
