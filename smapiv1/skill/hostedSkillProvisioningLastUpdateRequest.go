package skill

/*
HostedSkillProvisioningLastUpdateRequest Contains attributes related to last modification request of a hosted skill provisioning resource.
*/
type HostedSkillProvisioningLastUpdateRequest struct {
	Status   *Status              `json:"status"`
	Errors   []*StandardizedError `json:"errors"`
	Warnings []*StandardizedError `json:"warnings"`
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
