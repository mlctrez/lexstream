package catalog

/*
CatalogResponse CatalogId information.
*/
type CatalogResponse struct {
	// ID of the catalog created.
	CatalogId string `json:"catalogId,omitempty"`
}

/*
{
 "description": "CatalogId information.",
 "properties": {
  "catalogId": {
   "description": "ID of the catalog created.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
