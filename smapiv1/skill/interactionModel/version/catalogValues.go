package version

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
CatalogValues List of catalog values.
*/
type CatalogValues struct {
	IsTruncated bool   `json:"isTruncated"`
	NextToken   string `json:"nextToken"`
	// Total number of catalog values.
	TotalCount int            `json:"totalCount"`
	Links      *smapiv1.Links `json:"_links"`
	Values     []*ValueSchema `json:"values"`
}

/*
{
 "description": "List of catalog values.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "isTruncated": {
   "type": "boolean"
  },
  "nextToken": {
   "type": "string"
  },
  "totalCount": {
   "description": "Total number of catalog values.",
   "type": "integer"
  },
  "values": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.ValueSchema"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
