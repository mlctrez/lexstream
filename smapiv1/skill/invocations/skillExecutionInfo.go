package invocations

type SkillExecutionInfo struct {
	InvocationRequest  map[string]Request  `json,omitempty:"invocationRequest"`
	InvocationResponse map[string]Response `json,omitempty:"invocationResponse"`
	Metrics            map[string]Metrics  `json,omitempty:"metrics"`
}

/*
{
 "properties": {
  "invocationRequest": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.invocations.Request"
   },
   "type": "object"
  },
  "invocationResponse": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.invocations.Response"
   },
   "type": "object"
  },
  "metrics": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.invocations.Metrics"
   },
   "type": "object"
  }
 },
 "type": "object"
}
*/
