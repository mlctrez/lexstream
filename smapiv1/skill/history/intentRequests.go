package history

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
IntentRequests Response to the GET Intent Request History API. It contains the collection of utterances for the skill, nextToken and other metadata related to the search query.
*/
type IntentRequests struct {
	Links *smapiv1.Links `json,omitempty:"_links"`
	// This token can be used to load the next page of the result.
	NextToken string `json,omitempty:"nextToken"`
	// This property is true when all the results matching the search request haven't been returned, false otherwise.
	IsTruncated bool `json,omitempty:"isTruncated"`
	// Total number of records that matched the given search query.
	TotalCount int `json,omitempty:"totalCount"`
	// Position of the current page in the result set.
	StartIndex int `json,omitempty:"startIndex"`
	// The Skill Id.
	SkillId string `json,omitempty:"skillId"`
	// List of intent requests for the skill
	Items []*IntentRequest `json,omitempty:"items"`
}

/*
{
 "description": "Response to the GET Intent Request History API. It contains the collection of utterances for the skill, nextToken and other metadata related to the search query.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "isTruncated": {
   "description": "This property is true when all the results matching the search request haven't been returned, false otherwise.",
   "type": "boolean"
  },
  "items": {
   "description": "List of intent requests for the skill",
   "items": {
    "$ref": "#/definitions/v1.skill.history.IntentRequest"
   },
   "type": "array"
  },
  "nextToken": {
   "description": "This token can be used to load the next page of the result.",
   "type": "string"
  },
  "skillId": {
   "description": "The Skill Id.",
   "type": "string"
  },
  "startIndex": {
   "description": "Position of the current page in the result set.",
   "type": "number"
  },
  "totalCount": {
   "description": "Total number of records that matched the given search query.",
   "type": "number"
  }
 },
 "type": "object"
}
*/
