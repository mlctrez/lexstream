package isp

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListInSkillProduct List of in-skill products.
*/
type ListInSkillProduct struct {
	Links *smapiv1.Links `json:"_links,omitempty"`
	// Information for each in-skill product.
	InSkillProducts []*InSkillProductSummary `json:"inSkillProducts,omitempty"`
	IsTruncated     bool                     `json:"isTruncated,omitempty"`
	NextToken       string                   `json:"nextToken,omitempty"`
}

/*
{
 "description": "List of in-skill products.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "inSkillProducts": {
   "description": "Information for each in-skill product.",
   "items": {
    "$ref": "#/definitions/v1.isp.InSkillProductSummary"
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
