package skill

/*
HostedSkillDeploymentStatusLastUpdateRequest Contains attributes related to last modification request of a hosted skill deployment resource.
*/
type HostedSkillDeploymentStatusLastUpdateRequest struct {
	Status            *Status                       `json:"status"`
	Errors            []*StandardizedError          `json:"errors"`
	Warnings          []*StandardizedError          `json:"warnings"`
	DeploymentDetails *HostedSkillDeploymentDetails `json:"deploymentDetails"`
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
