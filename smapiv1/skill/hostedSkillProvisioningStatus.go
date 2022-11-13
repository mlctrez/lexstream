package skill

/*
HostedSkillProvisioningStatus Defines the provisioning status for hosted skill.
*/
type HostedSkillProvisioningStatus struct {
	LastUpdateRequest *HostedSkillProvisioningLastUpdateRequest `json:"lastUpdateRequest"`
}

/*
{
 "description": "Defines the provisioning status for hosted skill.",
 "properties": {
  "lastUpdateRequest": {
   "$ref": "#/definitions/v1.skill.HostedSkillProvisioningLastUpdateRequest"
  }
 },
 "type": "object"
}
*/
