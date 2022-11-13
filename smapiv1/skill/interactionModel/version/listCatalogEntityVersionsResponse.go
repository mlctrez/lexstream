package version

/*
ListCatalogEntityVersionsResponse List of catalog versions of a catalog for the vendor in sortDirection order, descending as default.
*/
type ListCatalogEntityVersionsResponse struct {
	Links *Links `json:"_links"`
	/*
	   List of catalog entity versions.
	*/
	CatalogVersions []*CatalogEntityVersion `json:"catalogVersions"`
	IsTruncated     bool                    `json:"isTruncated"`
	NextToken       string                  `json:"nextToken"`
	TotalCount      int                     `json:"totalCount"`
}

/*
{
 "description": "List of catalog versions of a catalog for the vendor in sortDirection order, descending as default.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.Links"
  },
  "catalogVersions": {
   "description": "List of catalog entity versions.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.CatalogEntityVersion"
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
