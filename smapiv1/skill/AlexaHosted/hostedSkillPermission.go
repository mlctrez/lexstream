package alexahosted

/*
HostedSkillPermission Customer's permission about Hosted skill features.
*/
type HostedSkillPermission struct {
	Permission *HostedSkillPermissionType   `json:"permission"`
	Status     *HostedSkillPermissionStatus `json:"status"`
	ActionUrl  string                       `json:"actionUrl"`
}

/*
{
 "description": "Customer's permission about Hosted skill features.",
 "properties": {
  "actionUrl": {
   "type": "string"
  },
  "permission": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillPermissionType",
   "x-isEnum": true
  },
  "status": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostedSkillPermissionStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
