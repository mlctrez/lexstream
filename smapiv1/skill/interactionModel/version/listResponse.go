package version

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListResponse List of interactionModel versions of a skill for the vendor
*/
type ListResponse struct {
	Links *smapiv1.Links `json:"_links,omitempty"`
	/*
	   List of interaction model versions.
	*/
	SkillModelVersions []*VersionItems `json:"skillModelVersions,omitempty"`
	IsTruncated        bool            `json:"isTruncated,omitempty"`
	NextToken          string          `json:"nextToken,omitempty"`
}

/*
{
 "description": "List of interactionModel versions of a skill for the vendor",
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
  "skillModelVersions": {
   "description": "List of interaction model versions.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.version.VersionItems"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
