package invocations

type InvokeSkillResponse struct {
	Status *InvocationResponseStatus `json,omitempty:"status"`
	Result *InvocationResponseResult `json,omitempty:"result"`
}

/*
{
 "properties": {
  "result": {
   "$ref": "#/definitions/v1.skill.invocations.InvocationResponseResult"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.invocations.InvocationResponseStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
