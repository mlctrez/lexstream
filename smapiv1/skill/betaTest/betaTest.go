package betatest

import "time"

/*
BetaTest Beta test for an Alexa skill.
*/
type BetaTest struct {
	// Expiry date of the beta test.
	ExpiryDate time.Time `json:"expiryDate"`
	Status     *Status   `json:"status"`
	// Email address for providing feedback
	FeedbackEmail string `json:"feedbackEmail"`
	// Deeplinking for getting authorized to the beta test.
	InvitationUrl string `json:"invitationUrl"`
	// Remaining invite count for the beta test.
	InvitesRemaining int `json:"invitesRemaining"`
}

/*
{
 "description": "Beta test for an Alexa skill.",
 "properties": {
  "expiryDate": {
   "description": "Expiry date of the beta test.",
   "format": "date-time",
   "type": "string"
  },
  "feedbackEmail": {
   "description": "Email address for providing feedback",
   "type": "string"
  },
  "invitationUrl": {
   "description": "Deeplinking for getting authorized to the beta test.",
   "type": "string"
  },
  "invitesRemaining": {
   "description": "Remaining invite count for the beta test.",
   "type": "number"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.betaTest.Status",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
