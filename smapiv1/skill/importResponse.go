package skill

type ImportResponse struct {
	Status   *ResponseStatus      `json,omitempty:"status"`
	Errors   []*StandardizedError `json,omitempty:"errors"`
	Warnings []*StandardizedError `json,omitempty:"warnings"`
	Skill    *ImportResponseSkill `json,omitempty:"skill"`
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
