package alexahosted

/*
HostedSkillPermission Customer's permission about Hosted skill features.
*/
type HostedSkillPermission struct {
	Permission *HostedSkillPermissionType   `json,omitempty:"permission"`
	Status     *HostedSkillPermissionStatus `json,omitempty:"status"`
	ActionUrl  string                       `json,omitempty:"actionUrl"`
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
