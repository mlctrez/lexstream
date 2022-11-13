package skill

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListSkillResponse List of skills for the vendor.
*/
type ListSkillResponse struct {
	Links *smapiv1.Links `json,omitempty:"_links"`
	/*
	   List of skill summaries. List might contain either one, two or three entries for a given skillId depending on the skill's publication history and the publication method.
	   `Skill containing certified stage`
	   * If a skill was never published to live, this list will contain two entries `:` one with stage 'development' and another with stage 'certified'. Both of these summaries will have same skillId.
	   * For any skill that has been published to 'live', this list will contain three entries `:` one with stage 'development', one with stage `certified` and one with stage 'live'. All of these summaries will have same skillId.
	   `Skill without certified stage`
	   * If a skill was never published to live, this list will contain only one entry for the skill with stage as 'development'.
	   * For any skill that has been published to 'live', this list will contain two entries `:` one with stage 'development' and another with stage 'live'. Both of these summaries will have same skillId.
	*/
	Skills      []*SkillSummary `json,omitempty:"skills"`
	IsTruncated bool            `json,omitempty:"isTruncated"`
	NextToken   string          `json,omitempty:"nextToken"`
}

/*
{
 "description": "List of skills for the vendor.",
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
  "skills": {
   "description": "List of skill summaries. List might contain either one, two or three entries for a given skillId depending on the skill's publication history and the publication method.\n`Skill containing certified stage`\n* If a skill was never published to live, this list will contain two entries `:` one with stage 'development' and another with stage 'certified'. Both of these summaries will have same skillId.\n* For any skill that has been published to 'live', this list will contain three entries `:` one with stage 'development', one with stage `certified` and one with stage 'live'. All of these summaries will have same skillId.\n`Skill without certified stage`\n* If a skill was never published to live, this list will contain only one entry for the skill with stage as 'development'.\n* For any skill that has been published to 'live', this list will contain two entries `:` one with stage 'development' and another with stage 'live'. Both of these summaries will have same skillId.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.SkillSummary"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
