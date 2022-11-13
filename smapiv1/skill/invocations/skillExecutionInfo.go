package invocations

type SkillExecutionInfo struct {
	InvocationRequest  map[string]Request  `json:"invocationRequest,omitempty"`
	InvocationResponse map[string]Response `json:"invocationResponse,omitempty"`
	Metrics            map[string]Metrics  `json:"metrics,omitempty"`
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
