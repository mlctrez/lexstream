package manifest

type LocalizedFlashBriefingInfoItems struct {
	// Logical name of the feed. This is used to signify relation among feeds across different locales. Example If you have "weather" feed in multiple locale then consider naming it "weather_update" and we will make sure to play the right feed if customer changes the language on device.
	LogicalName string `json:"logicalName"`
	// Name that identifies this feed.
	Name string `json:"name"`
	// Url for the feed
	Url string `json:"url"`
	// Uri for the feed image
	ImageUri        string                        `json:"imageUri"`
	ContentType     *FlashBriefingContentType     `json:"contentType"`
	Genre           *FlashBriefingGenre           `json:"genre"`
	UpdateFrequency *FlashBriefingUpdateFrequency `json:"updateFrequency"`
	// A short introduction for the feed that Alexa reads to the customer before the feed contents. Should start with "In" or "From".
	VuiPreamble string `json:"vuiPreamble"`
	// True if this should be the default feed to be enabled when customer enables the skill false otherwise.
	IsDefault bool `json:"isDefault"`
}

/*
{
 "properties": {
  "contentType": {
   "$ref": "#/definitions/v1.skill.Manifest.FlashBriefingContentType",
   "x-isEnum": true
  },
  "genre": {
   "$ref": "#/definitions/v1.skill.Manifest.FlashBriefingGenre",
   "x-isEnum": true
  },
  "imageUri": {
   "description": "Uri for the feed image",
   "type": "string"
  },
  "isDefault": {
   "description": "True if this should be the default feed to be enabled when customer enables the skill false otherwise.",
   "type": "boolean"
  },
  "logicalName": {
   "description": "Logical name of the feed. This is used to signify relation among feeds across different locales. Example If you have \"weather\" feed in multiple locale then consider naming it \"weather_update\" and we will make sure to play the right feed if customer changes the language on device.",
   "type": "string"
  },
  "name": {
   "description": "Name that identifies this feed.",
   "type": "string"
  },
  "updateFrequency": {
   "$ref": "#/definitions/v1.skill.Manifest.FlashBriefingUpdateFrequency",
   "x-isEnum": true
  },
  "url": {
   "description": "Url for the feed",
   "type": "string"
  },
  "vuiPreamble": {
   "description": "A short introduction for the feed that Alexa reads to the customer before the feed contents. Should start with \"In\" or \"From\".",
   "type": "string"
  }
 },
 "type": "object"
}
*/
