package skill

/*
HostedSkillDeploymentStatusLastUpdateRequest Contains attributes related to last modification request of a hosted skill deployment resource.
*/
type HostedSkillDeploymentStatusLastUpdateRequest struct {
	Status            *Status                       `json:"status,omitempty"`
	Errors            []*StandardizedError          `json:"errors,omitempty"`
	Warnings          []*StandardizedError          `json:"warnings,omitempty"`
	DeploymentDetails *HostedSkillDeploymentDetails `json:"deploymentDetails,omitempty"`
}

/*
{
 "description": "Contains attributes related to last modification request of a hosted skill deployment resource.",
 "properties": {
  "deploymentDetails": {
   "$ref": "#/definitions/v1.skill.HostedSkillDeploymentDetails"
  },
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.Status",
   "x-isEnum": true
  },
  "warnings": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
