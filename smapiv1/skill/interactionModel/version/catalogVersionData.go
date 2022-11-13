package version

/*
CatalogVersionData Catalog version data with metadata.
*/
type CatalogVersionData struct {
	Source *InputSource `json,omitempty:"source"`
	// Description string for specific catalog version.
	Description string `json,omitempty:"description"`
	// Specific catalog version.
	Version string `json,omitempty:"version"`
}

/*
{
 "description": "Catalog version data with metadata.",
 "properties": {
  "description": {
   "description": "Description string for specific catalog version.",
   "type": "string"
  },
  "source": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.InputSource"
  },
  "version": {
   "description": "Specific catalog version.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
