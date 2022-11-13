package skill

type ExportResponse struct {
	Status *ResponseStatus      `json:"status,omitempty"`
	Skill  *ExportResponseSkill `json:"skill,omitempty"`
}

/*
{
 "properties": {
  "skill": {
   "$ref": "#/definitions/v1.skill.ExportResponseSkill"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.ResponseStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
