package version

/*
VersionData Catalog version specific data.
*/
type VersionData struct {
	Source *InputSource `json,omitempty:"source"`
	// Description string for specific catalog version.
	Description string `json,omitempty:"description"`
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
