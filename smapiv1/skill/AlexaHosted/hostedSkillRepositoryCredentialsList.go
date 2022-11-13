package alexahosted

/*
HostedSkillRepositoryCredentialsList defines the structure for the hosted skill repository credentials response
*/
type HostedSkillRepositoryCredentialsList struct {
	RepositoryCredentials *HostedSkillRepositoryCredentials `json:"repositoryCredentials"`
}

/*
{
 "description": "defines the structure for the hosted skill repository credentials response",
 "properties": {
  "repositoryCredentials": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillRepositoryCredentials"
  }
 },
 "type": "object"
}
*/
