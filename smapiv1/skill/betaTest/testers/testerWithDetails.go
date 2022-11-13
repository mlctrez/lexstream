package testers

import "time"

/*
TesterWithDetails Tester information.
*/
type TesterWithDetails struct {
	// Email address of the tester.
	EmailId string `json,omitempty:"emailId"`
	// Date and time when the tester is added to the beta test.
	AssociationDate time.Time `json,omitempty:"associationDate"`
	// Indicates whether the tester is allowed to be sent reminder.
	IsReminderAllowed bool              `json,omitempty:"isReminderAllowed"`
	InvitationStatus  *InvitationStatus `json,omitempty:"invitationStatus"`
}

/*
{
 "description": "Tester information.",
 "properties": {
  "associationDate": {
   "description": "Date and time when the tester is added to the beta test.",
   "format": "date-time",
   "type": "string"
  },
  "emailId": {
   "description": "Email address of the tester.",
   "type": "string"
  },
  "invitationStatus": {
   "$ref": "#/definitions/v1.skill.betaTest.testers.InvitationStatus",
   "x-isEnum": true
  },
  "isReminderAllowed": {
   "description": "Indicates whether the tester is allowed to be sent reminder.",
   "type": "boolean"
  }
 },
 "type": "object"
}
*/
