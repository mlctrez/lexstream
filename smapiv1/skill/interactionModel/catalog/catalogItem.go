package catalog

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
CatalogItem Definition for catalog entity.
*/
type CatalogItem struct {
	// Name of the catalog.
	Name string `json:"name"`
	// Description string about the catalog.
	Description string `json:"description"`
	// Identifier of the catalog, optional in get response as the request already has catalogId.
	CatalogId string         `json:"catalogId"`
	Links     *smapiv1.Links `json:"_links"`
}

/*
{
 "description": "Definition for catalog entity.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "catalogId": {
   "description": "Identifier of the catalog, optional in get response as the request already has catalogId.",
   "type": "string"
  },
  "description": {
   "description": "Description string about the catalog.",
   "type": "string"
  },
  "name": {
   "description": "Name of the catalog.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
