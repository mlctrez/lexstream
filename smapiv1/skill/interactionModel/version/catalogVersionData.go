package version

/*
CatalogVersionData Catalog version data with metadata.
*/
type CatalogVersionData struct {
	Source *InputSource `json:"source"`
	// Description string for specific catalog version.
	Description string `json:"description"`
	// Specific catalog version.
	Version string `json:"version"`
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
