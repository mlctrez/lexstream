package invocations

import skill "github.com/mlctrez/lexstream/smapiv1/skill"

type InvocationResponseResult struct {
	SkillExecutionInfo *SkillExecutionInfo      `json:"skillExecutionInfo,omitempty"`
	Error              *skill.StandardizedError `json:"error,omitempty"`
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
