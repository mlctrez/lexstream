package simulations

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type SimulationResult struct {
	AlexaExecutionInfo *AlexaExecutionInfo `json,omitempty:"alexaExecutionInfo"`
	SkillExecutionInfo *Invocation         `json,omitempty:"skillExecutionInfo"`
	Error              *smapiv1.Error      `json,omitempty:"error"`
}

/*
{
 "properties": {
  "alexaExecutionInfo": {
   "$ref": "#/definitions/v1.skill.simulations.AlexaExecutionInfo"
  },
  "error": {
   "$ref": "#/definitions/v1.Error"
  },
  "skillExecutionInfo": {
   "$ref": "#/definitions/v1.skill.simulations.Invocation"
  }
 },
 "type": "object"
}
*/
