package version

/*
CatalogUpdate Catalog update description object.
*/
type CatalogUpdate struct {
	// The catalog description with a 255 character maximum.
	Description string `json,omitempty:"description"`
}

/*
{
 "description": "Catalog update description object.",
 "properties": {
  "description": {
   "description": "The catalog description with a 255 character maximum.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
