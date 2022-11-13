package invocations

import skill "github.com/mlctrez/lexstream/smapiv1/skill"

type InvocationResponseResult struct {
	SkillExecutionInfo *SkillExecutionInfo      `json,omitempty:"skillExecutionInfo"`
	Error              *skill.StandardizedError `json,omitempty:"error"`
}

/*
{
 "properties": {
  "error": {
   "$ref": "#/definitions/v1.skill.StandardizedError"
  },
  "skillExecutionInfo": {
   "$ref": "#/definitions/v1.skill.invocations.SkillExecutionInfo"
  }
 },
 "type": "object"
}
*/
