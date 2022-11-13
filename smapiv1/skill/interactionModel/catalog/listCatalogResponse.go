package catalog

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListCatalogResponse List of catalog versions of a skill for the vendor.
*/
type ListCatalogResponse struct {
	Links *smapiv1.Links `json:"_links,omitempty"`
	/*
	   List of catalogs.
	*/
	Catalogs    []*CatalogItem `json:"catalogs,omitempty"`
	IsTruncated bool           `json:"isTruncated,omitempty"`
	NextToken   string         `json:"nextToken,omitempty"`
	TotalCount  int            `json:"totalCount,omitempty"`
}

/*
{
 "description": "List of catalog versions of a skill for the vendor.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "catalogs": {
   "description": "List of catalogs.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogItem"
   },
   "type": "array"
  },
  "isTruncated": {
   "type": "boolean"
  },
  "nextToken": {
   "type": "string"
  },
  "totalCount": {
   "type": "integer"
  }
 },
 "type": "object"
}
*/
