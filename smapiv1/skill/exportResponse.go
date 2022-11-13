package skill

type ExportResponse struct {
	Status *ResponseStatus      `json,omitempty:"status"`
	Skill  *ExportResponseSkill `json,omitempty:"skill"`
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
