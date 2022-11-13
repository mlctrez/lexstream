package skill

/*
BuildStep Describes the status of a build step.
*/
type BuildStep struct {
	Name   *BuildStepName       `json:"name,omitempty"`
	Status *Status              `json:"status,omitempty"`
	Errors []*StandardizedError `json:"errors,omitempty"`
}

/*
{
 "description": "Describes the status of a build step.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "name": {
   "$ref": "#/definitions/v1.skill.BuildStepName",
   "x-isEnum": true
  },
  "status": {
   "$ref": "#/definitions/v1.skill.Status",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
