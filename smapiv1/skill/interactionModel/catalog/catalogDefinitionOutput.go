package catalog

/*
CatalogDefinitionOutput Catalog request definitions.
*/
type CatalogDefinitionOutput struct {
	Catalog *CatalogEntity `json:"catalog"`
	// Time of the catalog definition creation.
	CreationTime string `json:"creationTime"`
	// Total number of versions.
	TotalVersions string `json:"totalVersions"`
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
