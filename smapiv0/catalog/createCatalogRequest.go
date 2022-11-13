package catalog

type CreateCatalogRequest struct {
	// Title of the catalog.
	Title string        `json:"title"`
	Type_ *CatalogType  `json:"type"`
	Usage *CatalogUsage `json:"usage"`
	// ID of the vendor owning the catalog.
	VendorId string `json:"vendorId"`
}

/*
{
 "properties": {
  "title": {
   "description": "Title of the catalog.",
   "type": "string"
  },
  "type": {
   "$ref": "#/definitions/v0.catalog.CatalogType",
   "x-isEnum": true
  },
  "usage": {
   "$ref": "#/definitions/v0.catalog.CatalogUsage",
   "x-isEnum": true
  },
  "vendorId": {
   "description": "ID of the vendor owning the catalog.",
   "type": "string"
  }
 },
 "required": [
  "title",
  "type",
  "usage",
  "vendorId"
 ],
 "type": "object"
}
*/