package manifest

/*
SkillManifestPublishingInformation Defines the structure for publishing information in the skill manifest.
*/
type SkillManifestPublishingInformation struct {
	// Name of the skill that is displayed to customers in the Alexa app.
	Name string `json:"name,omitempty"`
	// Description of the skill's purpose and feature and how it works. Should describe any prerequisites like hardware or account requirements and detailed steps for the customer to get started. For Flash Briefing skill list the feeds offered within the skill. Use a conversational tone and correct grammar and punctuation. This description displays to customers on the skill detail card in the Alexa app.
	Description string `json:"description,omitempty"`
	// Defines the structure for locale specific publishing information in the skill manifest.
	Locales map[string]SkillManifestLocalizedPublishingInformation `json:"locales,omitempty"`
	// True if the skill should be distributed in all countries where Amazon distributes skill false otherwise.
	IsAvailableWorldwide bool                   `json:"isAvailableWorldwide,omitempty"`
	DistributionMode     *DistributionMode      `json:"distributionMode,omitempty"`
	GadgetSupport        *ManifestGadgetSupport `json:"gadgetSupport,omitempty"`
	// Special instructions provided by the developer to test the skill.
	TestingInstructions string `json:"testingInstructions,omitempty"`
	// Category that best describes a skill. Indicates the filter category for the skill in the Alexa App.
	Category string `json:"category,omitempty"`
	// Selected list of countries provided by the skill owner where Amazon can distribute the skill.
	DistributionCountries []*DistributionCountries `json:"distributionCountries,omitempty"`
}

/*
{
 "description": "Defines the structure for publishing information in the skill manifest.",
 "properties": {
  "category": {
   "description": "Category that best describes a skill. Indicates the filter category for the skill in the Alexa App.",
   "type": "string"
  },
  "description": {
   "description": "Description of the skill's purpose and feature and how it works. Should describe any prerequisites like hardware or account requirements and detailed steps for the customer to get started. For Flash Briefing skill list the feeds offered within the skill. Use a conversational tone and correct grammar and punctuation. This description displays to customers on the skill detail card in the Alexa app.",
   "type": "string"
  },
  "distributionCountries": {
   "description": "Selected list of countries provided by the skill owner where Amazon can distribute the skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.DistributionCountries"
   },
   "type": "array"
  },
  "distributionMode": {
   "$ref": "#/definitions/v1.skill.Manifest.DistributionMode",
   "x-isEnum": true
  },
  "gadgetSupport": {
   "$ref": "#/definitions/v1.skill.Manifest.ManifestGadgetSupport"
  },
  "isAvailableWorldwide": {
   "description": "True if the skill should be distributed in all countries where Amazon distributes skill false otherwise.",
   "type": "boolean"
  },
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.SkillManifestLocalizedPublishingInformation"
   },
   "description": "Defines the structure for locale specific publishing information in the skill manifest.",
   "type": "object"
  },
  "name": {
   "description": "Name of the skill that is displayed to customers in the Alexa app.",
   "type": "string"
  },
  "testingInstructions": {
   "description": "Special instructions provided by the developer to test the skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
