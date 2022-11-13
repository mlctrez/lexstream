package testers

/*
TestersList List of testers.
*/
type TestersList struct {
	// List of the email address of testers.
	Testers []*Tester `json:"testers,omitempty"`
}

/*
{
 "description": "List of testers.",
 "properties": {
  "testers": {
   "description": "List of the email address of testers.",
   "items": {
    "$ref": "#/definitions/v1.skill.betaTest.testers.Tester"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
