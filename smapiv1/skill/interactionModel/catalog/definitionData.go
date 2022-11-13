package catalog

/*
DefinitionData Catalog request definitions.
*/
type DefinitionData struct {
	Catalog *CatalogInput `json:"catalog"`
	// The vendorId that the catalog should belong to.
	VendorId string `json:"vendorId"`
}

/*
{
 "description": "Catalog request definitions.",
 "properties": {
  "catalog": {
   "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogInput"
  },
  "vendorId": {
   "description": "The vendorId that the catalog should belong to.",
   "type": "string"
  }
 },
 "required": [
  "catalog",
  "vendorId"
 ],
 "type": "object"
}
*/