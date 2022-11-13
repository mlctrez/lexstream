package testers

/*
InvitationStatus Indicates whether the tester has accepted the invitation.
*/
type InvitationStatus string

func InvitationStatus_ACCEPTED() InvitationStatus {
	return "ACCEPTED"
}
func InvitationStatus_NOT_ACCEPTED() InvitationStatus {
	return "NOT_ACCEPTED"
}
