package skill

type ImportResponse struct {
	Status   *ResponseStatus      `json:"status,omitempty"`
	Errors   []*StandardizedError `json:"errors,omitempty"`
	Warnings []*StandardizedError `json:"warnings,omitempty"`
	Skill    *ImportResponseSkill `json:"skill,omitempty"`
}

/*
{
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "skill": {
   "$ref": "#/definitions/v1.skill.ImportResponseSkill"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.ResponseStatus",
   "x-isEnum": true
  },
  "warnings": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
