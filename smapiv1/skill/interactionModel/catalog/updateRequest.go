package catalog

/*
UpdateRequest Catalog update request object.
*/
type UpdateRequest struct {
	// The catalog name.
	Name string `json:"name,omitempty"`
	// The catalog description with a 255 character maximum.
	Description string `json:"description,omitempty"`
}

/*
{
 "description": "Catalog update request object.",
 "properties": {
  "description": {
   "description": "The catalog description with a 255 character maximum.",
   "type": "string"
  },
  "name": {
   "description": "The catalog name.",
   "type": "string"
  }
 },
 "required": [
  "description",
  "name"
 ],
 "type": "object"
}
*/
