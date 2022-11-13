package evaluations

type Expected struct {
	Domain string          `json,omitempty:"domain"`
	Intent *ExpectedIntent `json,omitempty:"intent"`
}

/*
{
 "properties": {
  "domain": {
   "type": "string"
  },
  "intent": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.ExpectedIntent"
  }
 },
 "type": "object"
}
*/
