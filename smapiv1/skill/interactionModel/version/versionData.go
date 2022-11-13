package version

/*
VersionData Catalog version specific data.
*/
type VersionData struct {
	Source *InputSource `json:"source,omitempty"`
	// Description string for specific catalog version.
	Description string `json:"description,omitempty"`
}

/*
{
 "description": "Catalog version specific data.",
 "properties": {
  "description": {
   "description": "Description string for specific catalog version.",
   "type": "string"
  },
  "source": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.InputSource"
  }
 },
 "type": "object"
}
*/
