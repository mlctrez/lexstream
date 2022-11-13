package betatest

/*
TestBody Beta test meta-data.
*/
type TestBody struct {
	// Email address for providing feedback.
	FeedbackEmail string `json:"feedbackEmail,omitempty"`
}

/*
{
 "description": "Beta test meta-data.",
 "properties": {
  "feedbackEmail": {
   "description": "Email address for providing feedback.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
