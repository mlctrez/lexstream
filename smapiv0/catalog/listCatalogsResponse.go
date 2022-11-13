package catalog

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

/*
ListCatalogsResponse Information about catalogs.
*/
type ListCatalogsResponse struct {
	Links *smapiv0.Links `json,omitempty:"_links"`
	/*
	   List of catalog summaries.
	*/
	Catalogs    []*CatalogSummary `json,omitempty:"catalogs"`
	IsTruncated bool              `json,omitempty:"isTruncated"`
	NextToken   string            `json,omitempty:"nextToken"`
}

/*
{
 "description": "Information about catalogs.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v0.Links"
  },
  "catalogs": {
   "description": "List of catalog summaries.\n",
   "items": {
    "$ref": "#/definitions/v0.catalog.CatalogSummary"
   },
   "type": "array"
  },
  "isTruncated": {
   "type": "boolean"
  },
  "nextToken": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
