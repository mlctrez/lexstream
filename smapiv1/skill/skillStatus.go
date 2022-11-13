package skill

/*
SkillStatus Defines the structure for skill status response.
*/
type SkillStatus struct {
	Manifest *ManifestStatus `json:"manifest,omitempty"`
	// Status for available interaction models, keyed by locale.
	InteractionModel        map[string]SkillInteractionModelStatus `json:"interactionModel,omitempty"`
	HostedSkillDeployment   *HostedSkillDeploymentStatus           `json:"hostedSkillDeployment,omitempty"`
	HostedSkillProvisioning *HostedSkillProvisioningStatus         `json:"hostedSkillProvisioning,omitempty"`
}

/*
{
 "description": "Defines the structure for skill status response.",
 "properties": {
  "hostedSkillDeployment": {
   "$ref": "#/definitions/v1.skill.HostedSkillDeploymentStatus"
  },
  "hostedSkillProvisioning": {
   "$ref": "#/definitions/v1.skill.HostedSkillProvisioningStatus"
  },
  "interactionModel": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.SkillInteractionModelStatus"
   },
   "description": "Status for available interaction models, keyed by locale.",
   "type": "object"
  },
  "manifest": {
   "$ref": "#/definitions/v1.skill.ManifestStatus"
  }
 },
 "type": "object"
}
*/
