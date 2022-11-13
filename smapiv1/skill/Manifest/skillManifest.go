package manifest

/*
SkillManifest Defines the structure for a skill's metadata.
*/
type SkillManifest struct {
	// Version of the skill manifest.
	ManifestVersion       string                              `json,omitempty:"manifestVersion"`
	PublishingInformation *SkillManifestPublishingInformation `json,omitempty:"publishingInformation"`
	PrivacyAndCompliance  *SkillManifestPrivacyAndCompliance  `json,omitempty:"privacyAndCompliance"`
	Events                *SkillManifestEvents                `json,omitempty:"events"`
	// Defines the structure for required permissions information in the skill manifest.
	Permissions []*PermissionItems `json,omitempty:"permissions"`
	Apis        *SkillManifestApis `json,omitempty:"apis"`
}

/*
{
 "description": "Defines the structure for a skill's metadata.",
 "properties": {
  "apis": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestApis"
  },
  "events": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEvents"
  },
  "manifestVersion": {
   "description": "Version of the skill manifest.",
   "type": "string"
  },
  "permissions": {
   "description": "Defines the structure for required permissions information in the skill manifest.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.PermissionItems"
   },
   "type": "array"
  },
  "privacyAndCompliance": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestPrivacyAndCompliance"
  },
  "publishingInformation": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestPublishingInformation"
  }
 },
 "type": "object"
}
*/
