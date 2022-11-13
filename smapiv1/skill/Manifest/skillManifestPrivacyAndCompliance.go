package manifest

/*
SkillManifestPrivacyAndCompliance Defines the structure for privacy & compliance information in the skill manifest.
*/
type SkillManifestPrivacyAndCompliance struct {
	// Defines the structure for locale specific privacy & compliance information in the skill manifest.
	Locales map[string]SkillManifestLocalizedPrivacyAndCompliance `json:"locales,omitempty"`
	// True if the skill allows users to make purchases or spend real money false otherwise.
	AllowsPurchases bool `json:"allowsPurchases,omitempty"`
	// True if the skill collects users' personal information false otherwise.
	UsesPersonalInfo bool `json:"usesPersonalInfo,omitempty"`
	// True if the skill is directed to or targets children under the age of 13/16 false otherwise.
	IsChildDirected bool `json:"isChildDirected,omitempty"`
	// True if it is certified that the skill may be imported to and exported from the United States and all other countries and regions in which Amazon operate its program or in which skill owner have authorized sales to end users (without the need for Amazon to obtain any license or clearance or take any other action) and is in full compliance with all applicable laws and regulations governing imports and export including those applicable to software that makes use of encryption technology.
	IsExportCompliant bool `json:"isExportCompliant,omitempty"`
	// True if the skill contains advertising false otherwise.
	ContainsAds bool `json:"containsAds,omitempty"`
	// True if the skill developer is a Covered Entity (CE) or Business Associate (BA) as defined by the Health Insurance Portability And Accountability Act (HIPAA) and the skill requires Amazon to process PHI on their behalf, false otherwise. This is an optional property and treated as false if not set.
	UsesHealthInfo bool `json:"usesHealthInfo,omitempty"`
}

/*
{
 "description": "Defines the structure for privacy \u0026 compliance information in the skill manifest.",
 "properties": {
  "allowsPurchases": {
   "description": "True if the skill allows users to make purchases or spend real money false otherwise.",
   "type": "boolean"
  },
  "containsAds": {
   "description": "True if the skill contains advertising false otherwise.",
   "type": "boolean"
  },
  "isChildDirected": {
   "description": "True if the skill is directed to or targets children under the age of 13/16 false otherwise.",
   "type": "boolean"
  },
  "isExportCompliant": {
   "description": "True if it is certified that the skill may be imported to and exported from the United States and all other countries and regions in which Amazon operate its program or in which skill owner have authorized sales to end users (without the need for Amazon to obtain any license or clearance or take any other action) and is in full compliance with all applicable laws and regulations governing imports and export including those applicable to software that makes use of encryption technology.",
   "type": "boolean"
  },
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.SkillManifestLocalizedPrivacyAndCompliance"
   },
   "description": "Defines the structure for locale specific privacy \u0026 compliance information in the skill manifest.",
   "type": "object"
  },
  "usesHealthInfo": {
   "description": "True if the skill developer is a Covered Entity (CE) or Business Associate (BA) as defined by the Health Insurance Portability And Accountability Act (HIPAA) and the skill requires Amazon to process PHI on their behalf, false otherwise. This is an optional property and treated as false if not set.",
   "type": "boolean"
  },
  "usesPersonalInfo": {
   "description": "True if the skill collects users' personal information false otherwise.",
   "type": "boolean"
  }
 },
 "type": "object"
}
*/
