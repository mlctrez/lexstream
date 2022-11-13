package manifest

/*
VideoApis Defines the structure for video api of the skill.
*/
type VideoApis struct {
	// Defines the structure for region information.
	Regions map[string]VideoRegion `json:"regions"`
	// Defines the structure for the locale specific video api information.
	Locales  map[string]VideoApisLocale `json:"locales"`
	Endpoint *LambdaEndpoint            `json:"endpoint"`
	// Object that contains <country> Objects for each supported country.
	Countries map[string]VideoCountryInfo `json:"countries"`
}

/*
{
 "description": "Defines the structure for video api of the skill.",
 "properties": {
  "countries": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.VideoCountryInfo"
   },
   "description": "Object that contains \u003ccountry\u003e Objects for each supported country.",
   "type": "object"
  },
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.LambdaEndpoint"
  },
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.VideoApisLocale"
   },
   "description": "Defines the structure for the locale specific video api information.",
   "type": "object"
  },
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.VideoRegion"
   },
   "description": "Defines the structure for region information.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
