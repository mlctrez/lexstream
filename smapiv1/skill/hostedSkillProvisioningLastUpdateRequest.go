package skill

/*
HostedSkillProvisioningLastUpdateRequest Contains attributes related to last modification request of a hosted skill provisioning resource.
*/
type HostedSkillProvisioningLastUpdateRequest struct {
	Status   *Status              `json:"status,omitempty"`
	Errors   []*StandardizedError `json:"errors,omitempty"`
	Warnings []*StandardizedError `json:"warnings,omitempty"`
}

/*
{
 "description": "Contains attributes related to last modification request of a hosted skill provisioning resource.",
 "properties": {
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
