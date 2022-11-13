package alexahosted

type HostedSkillPermissionStatus string

func HostedSkillPermissionStatus_ALLOWED() HostedSkillPermissionStatus {
	return "ALLOWED"
}
func HostedSkillPermissionStatus_NEW_USER_REGISTRATION_REQUIRED() HostedSkillPermissionStatus {
	return "NEW_USER_REGISTRATION_REQUIRED"
}
func HostedSkillPermissionStatus_RESOURCE_LIMIT_EXCEEDED() HostedSkillPermissionStatus {
	return "RESOURCE_LIMIT_EXCEEDED"
}
func HostedSkillPermissionStatus_RATE_EXCEEDED() HostedSkillPermissionStatus {
	return "RATE_EXCEEDED"
}
