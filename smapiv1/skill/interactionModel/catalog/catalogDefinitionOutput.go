package catalog

/*
CatalogDefinitionOutput Catalog request definitions.
*/
type CatalogDefinitionOutput struct {
	Catalog *CatalogEntity `json,omitempty:"catalog"`
	// Time of the catalog definition creation.
	CreationTime string `json,omitempty:"creationTime"`
	// Total number of versions.
	TotalVersions string `json,omitempty:"totalVersions"`
}

/*
{
 "description": "Catalog request definitions.",
 "properties": {
  "catalog": {
   "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogEntity"
  },
  "creationTime": {
   "description": "Time of the catalog definition creation.",
   "type": "string"
  },
  "totalVersions": {
   "description": "Total number of versions.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
