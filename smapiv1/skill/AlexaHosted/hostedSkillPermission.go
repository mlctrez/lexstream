package alexahosted

/*
HostedSkillPermission Customer's permission about Hosted skill features.
*/
type HostedSkillPermission struct {
	Permission *HostedSkillPermissionType   `json:"permission,omitempty"`
	Status     *HostedSkillPermissionStatus `json:"status,omitempty"`
	ActionUrl  string                       `json:"actionUrl,omitempty"`
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
