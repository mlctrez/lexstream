package skill

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListSkillVersionsResponse List of all skill versions
*/
type ListSkillVersionsResponse struct {
	Links *smapiv1.Links `json:"_links,omitempty"`
	// Skill version metadata
	SkillVersions []*SkillVersion `json:"skillVersions,omitempty"`
	IsTruncated   bool            `json:"isTruncated,omitempty"`
	NextToken     string          `json:"nextToken,omitempty"`
}

/*
{
 "description": "List of all skill versions",
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
  "skillVersions": {
   "description": "Skill version metadata",
   "items": {
    "$ref": "#/definitions/v1.skill.SkillVersion"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
