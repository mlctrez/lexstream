package version

/*
ListCatalogEntityVersionsResponse List of catalog versions of a catalog for the vendor in sortDirection order, descending as default.
*/
type ListCatalogEntityVersionsResponse struct {
	Links *Links `json:"_links,omitempty"`
	/*
	   List of catalog entity versions.
	*/
	CatalogVersions []*CatalogEntityVersion `json:"catalogVersions,omitempty"`
	IsTruncated     bool                    `json:"isTruncated,omitempty"`
	NextToken       string                  `json:"nextToken,omitempty"`
	TotalCount      int                     `json:"totalCount,omitempty"`
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
