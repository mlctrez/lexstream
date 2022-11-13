package interactionmodel

/*
ValueSupplier Supplier object to provide slot values.
*/
type ValueSupplier struct {
	// The exact type of validation e.g.CatalogValueSupplier etc.
	Type_ string `json:"type,omitempty"`
}

/*
{
 "description": "Supplier object to provide slot values.",
 "discriminator": "type",
 "properties": {
  "type": {
   "description": "The exact type of validation e.g.CatalogValueSupplier etc.",
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "required": [
  "type"
 ],
 "type": "object"
}
*/
