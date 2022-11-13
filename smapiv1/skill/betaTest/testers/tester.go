package testers

type Tester struct {
	// Email address of the tester.
	EmailId string `json,omitempty:"emailId"`
}

/*
{
 "properties": {
  "emailId": {
   "description": "Email address of the tester.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
