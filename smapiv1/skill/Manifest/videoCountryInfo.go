package manifest

/*
VideoCountryInfo Defines the structure of per-country video info in the skill manifest.
*/
type VideoCountryInfo struct {
	CatalogInformation []*VideoCatalogInfo `json,omitempty:"catalogInformation"`
}

/*
{
 "description": "Defines the structure of per-country video info in the skill manifest.",
 "properties": {
  "catalogInformation": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.VideoCatalogInfo"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
