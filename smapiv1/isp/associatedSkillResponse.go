package isp

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
AssociatedSkillResponse In-skill product skill association details.
*/
type AssociatedSkillResponse struct {
	// List of skill IDs that correspond to the skills associated with the in-skill product. The associations are stage specific. A live association is created through successful skill certification.
	AssociatedSkillIds []string       `json,omitempty:"associatedSkillIds"`
	Links              *smapiv1.Links `json,omitempty:"_links"`
	IsTruncated        bool           `json,omitempty:"isTruncated"`
	NextToken          string         `json,omitempty:"nextToken"`
}

/*
{
 "description": "In-skill product skill association details.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "associatedSkillIds": {
   "description": "List of skill IDs that correspond to the skills associated with the in-skill product. The associations are stage specific. A live association is created through successful skill certification.",
   "items": {
    "type": "string"
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
