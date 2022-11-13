package catalog

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

/*
ListCatalogsResponse Information about catalogs.
*/
type ListCatalogsResponse struct {
	Links *smapiv0.Links `json:"_links,omitempty"`
	/*
	   List of catalog summaries.
	*/
	Catalogs    []*CatalogSummary `json:"catalogs,omitempty"`
	IsTruncated bool              `json:"isTruncated,omitempty"`
	NextToken   string            `json:"nextToken,omitempty"`
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
