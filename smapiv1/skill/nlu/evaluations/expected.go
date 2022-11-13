package evaluations

type Expected struct {
	Domain string          `json:"domain"`
	Intent *ExpectedIntent `json:"intent"`
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
