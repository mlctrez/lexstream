package evaluations

type Expected struct {
	Domain string          `json:"domain,omitempty"`
	Intent *ExpectedIntent `json:"intent,omitempty"`
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
