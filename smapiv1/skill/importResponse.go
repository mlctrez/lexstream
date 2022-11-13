package skill

type ImportResponse struct {
	Status   *ResponseStatus      `json:"status"`
	Errors   []*StandardizedError `json:"errors"`
	Warnings []*StandardizedError `json:"warnings"`
	Skill    *ImportResponseSkill `json:"skill"`
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
