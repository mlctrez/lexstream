package isp

type UpdateInSkillProductRequest struct {
	InSkillProductDefinition *InSkillProductDefinition `json:"inSkillProductDefinition"`
}

/*
{
 "properties": {
  "inSkillProductDefinition": {
   "$ref": "#/definitions/v1.isp.InSkillProductDefinition"
  }
 },
 "type": "object"
}
*/
