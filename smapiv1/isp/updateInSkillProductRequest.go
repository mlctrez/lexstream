package isp

type UpdateInSkillProductRequest struct {
	InSkillProductDefinition *InSkillProductDefinition `json,omitempty:"inSkillProductDefinition"`
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
