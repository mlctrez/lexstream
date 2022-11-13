package simulations

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type SimulationResult struct {
	AlexaExecutionInfo *AlexaExecutionInfo `json:"alexaExecutionInfo,omitempty"`
	SkillExecutionInfo *Invocation         `json:"skillExecutionInfo,omitempty"`
	Error              *smapiv1.Error      `json:"error,omitempty"`
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
