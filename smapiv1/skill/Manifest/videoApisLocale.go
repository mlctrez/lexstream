package manifest

/*
VideoApisLocale Defines the structure for localized video api information.
*/
type VideoApisLocale struct {
	// Defines the video provider's targeting name.
	VideoProviderTargetingNames []string            `json:"videoProviderTargetingNames,omitempty"`
	VideoProviderLogoUri        string              `json:"videoProviderLogoUri,omitempty"`
	CatalogInformation          []*VideoCatalogInfo `json:"catalogInformation,omitempty"`
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
