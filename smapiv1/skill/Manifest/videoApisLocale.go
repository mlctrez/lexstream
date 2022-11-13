package manifest

/*
VideoApisLocale Defines the structure for localized video api information.
*/
type VideoApisLocale struct {
	// Defines the video provider's targeting name.
	VideoProviderTargetingNames []string            `json,omitempty:"videoProviderTargetingNames"`
	VideoProviderLogoUri        string              `json,omitempty:"videoProviderLogoUri"`
	CatalogInformation          []*VideoCatalogInfo `json,omitempty:"catalogInformation"`
}

/*
{
 "description": "Defines the structure for localized video api information.",
 "properties": {
  "catalogInformation": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.VideoCatalogInfo"
   },
   "type": "array"
  },
  "videoProviderLogoUri": {
   "type": "string"
  },
  "videoProviderTargetingNames": {
   "description": "Defines the video provider's targeting name.",
   "items": {
    "type": "string"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
