package manifest

/*
SkillManifestLocalizedPrivacyAndCompliance Defines the structure for locale specific privacy & compliance information in the skill manifest.
*/
type SkillManifestLocalizedPrivacyAndCompliance struct {
	// Link to the privacy policy that applies to this skill.
	PrivacyPolicyUrl string `json:"privacyPolicyUrl,omitempty"`
	// link to the terms of use document for this skill
	TermsOfUseUrl string `json:"termsOfUseUrl,omitempty"`
}

/*
{
 "description": "Defines the structure for locale specific privacy \u0026 compliance information in the skill manifest.",
 "properties": {
  "privacyPolicyUrl": {
   "description": "Link to the privacy policy that applies to this skill.",
   "type": "string"
  },
  "termsOfUseUrl": {
   "description": "link to the terms of use document for this skill",
   "type": "string"
  }
 },
 "type": "object"
}
*/
