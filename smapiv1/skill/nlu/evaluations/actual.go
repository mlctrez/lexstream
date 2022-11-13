package evaluations

type Actual struct {
	Domain string  `json,omitempty:"domain"`
	Intent *Intent `json,omitempty:"intent"`
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
