package catalog

/*
CatalogEntity Definition for catalog entity.
*/
type CatalogEntity struct {
	// Name of the catalog.
	Name string `json:"name"`
	// Description string about the catalog.
	Description string `json:"description"`
}

/*
{
 "description": "Definition for catalog entity.",
 "properties": {
  "description": {
   "description": "Description string about the catalog.",
   "type": "string"
  },
  "name": {
   "description": "Name of the catalog.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
