package evaluations

type Actual struct {
	Domain string  `json:"domain"`
	Intent *Intent `json:"intent"`
}

/*
{
 "properties": {
  "domain": {
   "type": "string"
  },
  "intent": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Intent"
  }
 },
 "type": "object"
}
*/
