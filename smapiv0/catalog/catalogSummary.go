package catalog

import "time"

type CatalogSummary struct {
	// Unique identifier of the added catalog object.
	Id string `json,omitempty:"id"`
	// Title of the catalog.
	Title string        `json,omitempty:"title"`
	Type_ *CatalogType  `json,omitempty:"type"`
	Usage *CatalogUsage `json,omitempty:"usage"`
	// The date time when the catalog was last updated.
	LastUpdatedDate time.Time `json,omitempty:"lastUpdatedDate"`
	// The date time when the catalog was created.
	CreatedDate time.Time `json,omitempty:"createdDate"`
	// The list of skill Ids associated with the catalog.
	AssociatedSkillIds []string `json,omitempty:"associatedSkillIds"`
}

/*
{
 "properties": {
  "associatedSkillIds": {
   "description": "The list of skill Ids associated with the catalog.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "createdDate": {
   "description": "The date time when the catalog was created.",
   "format": "date-time",
   "type": "string"
  },
  "id": {
   "description": "Unique identifier of the added catalog object.",
   "type": "string"
  },
  "lastUpdatedDate": {
   "description": "The date time when the catalog was last updated.",
   "format": "date-time",
   "type": "string"
  },
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
  }
 },
 "type": "object"
}
*/
