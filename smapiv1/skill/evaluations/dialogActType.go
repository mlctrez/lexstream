package evaluations

type DialogActType string

func DialogActType_DialogElicitSlot() DialogActType {
	return "Dialog.ElicitSlot"
}
func DialogActType_DialogConfirmSlot() DialogActType {
	return "Dialog.ConfirmSlot"
}
func DialogActType_DialogConfirmIntent() DialogActType {
	return "Dialog.ConfirmIntent"
}
