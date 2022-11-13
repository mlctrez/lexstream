package evaluations

/*
ConfirmationStatusType An enumeration indicating whether the user has explicitly confirmed or denied the entire intent. Possible values: "NONE", "CONFIRMED", "DENIED".
*/
type ConfirmationStatusType string

func ConfirmationStatusType_NONE() ConfirmationStatusType {
	return "NONE"
}
func ConfirmationStatusType_CONFIRMED() ConfirmationStatusType {
	return "CONFIRMED"
}
func ConfirmationStatusType_DENIED() ConfirmationStatusType {
	return "DENIED"
}
