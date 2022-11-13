package simulations

type Invocation struct {
	InvocationRequest  *InvocationRequest  `json,omitempty:"invocationRequest"`
	InvocationResponse *InvocationResponse `json,omitempty:"invocationResponse"`
	Metrics            *Metrics            `json,omitempty:"metrics"`
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
