package catalog

/*
CatalogInput Definition for catalog input.
*/
type CatalogInput struct {
	// Name of the catalog.
	Name string `json:"name"`
	// Description string about the catalog.
	Description string `json:"description"`
}

/*
{
 "description": "Definition for catalog input.",
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
 "required": [
  "name"
 ],
 "type": "object"
}
*/
