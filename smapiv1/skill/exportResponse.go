package skill

type ExportResponse struct {
	Status *ResponseStatus      `json:"status"`
	Skill  *ExportResponseSkill `json:"skill"`
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
