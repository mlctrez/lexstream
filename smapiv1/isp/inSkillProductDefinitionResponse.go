package isp

/*
InSkillProductDefinitionResponse Defines In-skill product response.
*/
type InSkillProductDefinitionResponse struct {
	InSkillProductDefinition *InSkillProductDefinition `json:"inSkillProductDefinition,omitempty"`
}

/*
{
 "description": "Defines In-skill product response.",
 "properties": {
  "inSkillProductDefinition": {
   "$ref": "#/definitions/v1.isp.InSkillProductDefinition"
  }
 },
 "type": "object"
}
*/
