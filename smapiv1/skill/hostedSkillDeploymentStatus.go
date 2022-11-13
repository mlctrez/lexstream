package skill

/*
HostedSkillDeploymentStatus Defines the most recent deployment status for the Alexa hosted skill.
*/
type HostedSkillDeploymentStatus struct {
	LastUpdateRequest *HostedSkillDeploymentStatusLastUpdateRequest `json:"lastUpdateRequest"`
}

/*
{
 "description": "Defines the most recent deployment status for the Alexa hosted skill.",
 "properties": {
  "lastUpdateRequest": {
   "$ref": "#/definitions/v1.skill.HostedSkillDeploymentStatusLastUpdateRequest"
  }
 },
 "type": "object"
}
*/
