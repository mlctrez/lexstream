package simulations

type Invocation struct {
	InvocationRequest  *InvocationRequest  `json:"invocationRequest"`
	InvocationResponse *InvocationResponse `json:"invocationResponse"`
	Metrics            *Metrics            `json:"metrics"`
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
