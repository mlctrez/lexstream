package isp

type UpdateInSkillProductRequest struct {
	InSkillProductDefinition *InSkillProductDefinition `json:"inSkillProductDefinition,omitempty"`
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
