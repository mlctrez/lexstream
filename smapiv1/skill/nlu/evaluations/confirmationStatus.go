package evaluations

/*
ConfirmationStatus An enumeration indicating whether the user has explicitly confirmed or denied the entire intent/slot. Possible values: 'NONE', 'CONFIRMED', 'DENIED'.
*/
type ConfirmationStatus string

func ConfirmationStatus_NONE() ConfirmationStatus {
	return "NONE"
}
func ConfirmationStatus_CONFIRMED() ConfirmationStatus {
	return "CONFIRMED"
}
func ConfirmationStatus_DENIED() ConfirmationStatus {
	return "DENIED"
}
