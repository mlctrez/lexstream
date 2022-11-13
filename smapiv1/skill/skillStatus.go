package skill

/*
SkillStatus Defines the structure for skill status response.
*/
type SkillStatus struct {
	Manifest *ManifestStatus `json,omitempty:"manifest"`
	// Status for available interaction models, keyed by locale.
	InteractionModel        map[string]SkillInteractionModelStatus `json,omitempty:"interactionModel"`
	HostedSkillDeployment   *HostedSkillDeploymentStatus           `json,omitempty:"hostedSkillDeployment"`
	HostedSkillProvisioning *HostedSkillProvisioningStatus         `json,omitempty:"hostedSkillProvisioning"`
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
