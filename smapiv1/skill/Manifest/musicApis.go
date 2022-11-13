package manifest

/*
MusicApis Defines the structure of music api in the skill manifest.
*/
type MusicApis struct {
	// Contains an array of the supported <region> Objects.
	Regions  map[string]LambdaRegion `json,omitempty:"regions"`
	Endpoint *LambdaEndpoint         `json,omitempty:"endpoint"`
	// Defines the structure of music capabilities information in the skill manifest.
	Capabilities []*MusicCapability `json,omitempty:"capabilities"`
	// A list of music skill interfaces that your skill supports.
	Interfaces []*MusicInterfaces `json,omitempty:"interfaces"`
	// Defines the structure of locale specific music information in the skill manifest.
	Locales map[string]LocalizedMusicInfo `json,omitempty:"locales"`
	// List of the type of content to be provided by the music skill.
	ContentTypes []*MusicContentType `json,omitempty:"contentTypes"`
}

/*
{
 "description": "Defines the structure of music api in the skill manifest.",
 "properties": {
  "capabilities": {
   "description": "Defines the structure of music capabilities information in the skill manifest.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicCapability"
   },
   "type": "array"
  },
  "contentTypes": {
   "description": "List of the type of content to be provided by the music skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicContentType"
   },
   "type": "array"
  },
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.LambdaEndpoint"
  },
  "interfaces": {
   "description": "A list of music skill interfaces that your skill supports.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicInterfaces"
   },
   "type": "array"
  },
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.LocalizedMusicInfo"
   },
   "description": "Defines the structure of locale specific music information in the skill manifest.",
   "type": "object"
  },
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.LambdaRegion"
   },
   "description": "Contains an array of the supported \u003cregion\u003e Objects.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
