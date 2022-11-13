package manifest

/*
LocalizedMusicInfo Defines the structure of localized music information in the skill manifest.
*/
type LocalizedMusicInfo struct {
	// Name to be used when Alexa renders the music skill name.
	PromptName    string           `json:"promptName"`
	Aliases       []*MusicAlias    `json:"aliases"`
	Features      []*MusicFeature  `json:"features"`
	WordmarkLogos []*MusicWordmark `json:"wordmarkLogos"`
}

/*
{
 "description": "Defines the structure of localized music information in the skill manifest.",
 "properties": {
  "aliases": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicAlias"
   },
   "type": "array"
  },
  "features": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicFeature"
   },
   "type": "array"
  },
  "promptName": {
   "description": "Name to be used when Alexa renders the music skill name.",
   "type": "string"
  },
  "wordmarkLogos": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicWordmark"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
