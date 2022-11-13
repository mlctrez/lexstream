package invocations

type InvokeSkillResponse struct {
	Status *InvocationResponseStatus `json:"status"`
	Result *InvocationResponseResult `json:"result"`
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
