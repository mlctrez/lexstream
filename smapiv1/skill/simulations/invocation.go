package simulations

type Invocation struct {
	InvocationRequest  *InvocationRequest  `json:"invocationRequest,omitempty"`
	InvocationResponse *InvocationResponse `json:"invocationResponse,omitempty"`
	Metrics            *Metrics            `json:"metrics,omitempty"`
}

/*
{
 "properties": {
  "invocationRequest": {
   "$ref": "#/definitions/v1.skill.simulations.InvocationRequest"
  },
  "invocationResponse": {
   "$ref": "#/definitions/v1.skill.simulations.InvocationResponse"
  },
  "metrics": {
   "$ref": "#/definitions/v1.skill.simulations.Metrics"
  }
 },
 "type": "object"
}
*/
