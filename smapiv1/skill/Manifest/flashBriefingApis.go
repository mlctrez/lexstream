package manifest

/*
FlashBriefingApis Defines the structure for flash briefing api of the skill.
*/
type FlashBriefingApis struct {
	// Defines the structure for locale specific flash briefing api information.
	Locales map[string]LocalizedFlashBriefingInfo `json,omitempty:"locales"`
}

/*
{
 "description": "Defines the structure for flash briefing api of the skill.",
 "properties": {
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.LocalizedFlashBriefingInfo"
   },
   "description": "Defines the structure for locale specific flash briefing api information.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
