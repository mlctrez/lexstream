package alexahosted

type HostedSkillRepositoryCredentialsRequest struct {
	Repository *HostedSkillRepositoryInfo `json:"repository,omitempty"`
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
