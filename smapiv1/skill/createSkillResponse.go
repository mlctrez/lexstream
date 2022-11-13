package skill

/*
CreateSkillResponse SkillId information.
*/
type CreateSkillResponse struct {
	// ID of the skill created.
	SkillId string `json:"skillId,omitempty"`
}

/*
{
 "description": "SkillId information.",
 "properties": {
  "skillId": {
   "description": "ID of the skill created.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
