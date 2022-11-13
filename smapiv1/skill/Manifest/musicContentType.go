package manifest

/*
MusicContentType Defines the structure for content that can be provided by a music skill.
*/
type MusicContentType struct {
	Name *MusicContentName `json:"name,omitempty"`
}

/*
{
 "description": "Defines the structure for content that can be provided by a music skill.",
 "properties": {
  "name": {
   "$ref": "#/definitions/v1.skill.Manifest.MusicContentName",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
