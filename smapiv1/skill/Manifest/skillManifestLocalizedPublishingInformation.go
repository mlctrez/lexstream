package manifest

/*
SkillManifestLocalizedPublishingInformation Defines the structure for locale specific publishing information in the skill manifest.
*/
type SkillManifestLocalizedPublishingInformation struct {
	// Name of the skill that is displayed to customers in the Alexa app.
	Name string `json,omitempty:"name"`
	// URL to a small icon for the skill, which is shown in the list of skills (108x108px).
	SmallIconUri string `json,omitempty:"smallIconUri"`
	// URL to a large icon that represents this skill (512x512px).
	LargeIconUri string `json,omitempty:"largeIconUri"`
	// Summary description of the skill, which is shown when viewing the list of skills.
	Summary string `json,omitempty:"summary"`
	// A full description explaining the skill’s core functionality and any prerequisites to using it (such as additional hardware, software, or accounts). For a Flash Briefing skill, you must list the feeds for the skill.
	Description string `json,omitempty:"description"`
	// Updates description of the skill's new features and fixes in the version. Should describe changes in the revisions of the skill.
	UpdatesDescription string `json,omitempty:"updatesDescription"`
	// Three example phrases that illustrate how users can invoke your skill. For accuracy, these phrases must come directly from your sample utterances.
	ExamplePhrases []string `json,omitempty:"examplePhrases"`
	// Sample keyword phrases that describe the skill.
	Keywords []string `json,omitempty:"keywords"`
}

/*
{
 "description": "Defines the structure for locale specific publishing information in the skill manifest.",
 "properties": {
  "description": {
   "description": "A full description explaining the skill’s core functionality and any prerequisites to using it (such as additional hardware, software, or accounts). For a Flash Briefing skill, you must list the feeds for the skill.",
   "type": "string"
  },
  "examplePhrases": {
   "description": "Three example phrases that illustrate how users can invoke your skill. For accuracy, these phrases must come directly from your sample utterances.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "keywords": {
   "description": "Sample keyword phrases that describe the skill.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "largeIconUri": {
   "description": "URL to a large icon that represents this skill (512x512px).",
   "type": "string"
  },
  "name": {
   "description": "Name of the skill that is displayed to customers in the Alexa app.",
   "type": "string"
  },
  "smallIconUri": {
   "description": "URL to a small icon for the skill, which is shown in the list of skills (108x108px).",
   "type": "string"
  },
  "summary": {
   "description": "Summary description of the skill, which is shown when viewing the list of skills.",
   "type": "string"
  },
  "updatesDescription": {
   "description": "Updates description of the skill's new features and fixes in the version. Should describe changes in the revisions of the skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
