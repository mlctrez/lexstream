package evaluations

type Actual struct {
	Domain string  `json:"domain,omitempty"`
	Intent *Intent `json:"intent,omitempty"`
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
