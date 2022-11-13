package catalog

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListCatalogResponse List of catalog versions of a skill for the vendor.
*/
type ListCatalogResponse struct {
	Links *smapiv1.Links `json,omitempty:"_links"`
	/*
	   List of catalogs.
	*/
	Catalogs    []*CatalogItem `json,omitempty:"catalogs"`
	IsTruncated bool           `json,omitempty:"isTruncated"`
	NextToken   string         `json,omitempty:"nextToken"`
	TotalCount  int            `json,omitempty:"totalCount"`
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
