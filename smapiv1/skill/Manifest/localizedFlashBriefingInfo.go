package manifest

/*
LocalizedFlashBriefingInfo Defines the localized flash briefing api information.
*/
type LocalizedFlashBriefingInfo struct {
	// Defines the structure for a feed information in the skill manifest.
	Feeds []*LocalizedFlashBriefingInfoItems `json:"feeds,omitempty"`
	// Alexa says this to the customer if the skill fails to render the content.
	CustomErrorMessage string `json:"customErrorMessage,omitempty"`
}

/*
{
 "description": "Defines the localized flash briefing api information.",
 "properties": {
  "customErrorMessage": {
   "description": "Alexa says this to the customer if the skill fails to render the content.",
   "type": "string"
  },
  "feeds": {
   "description": "Defines the structure for a feed information in the skill manifest.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.LocalizedFlashBriefingInfoItems"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
