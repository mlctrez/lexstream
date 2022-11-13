package history

/*
DialogActName Dialog act directive name.
* `Dialog.ElicitSlot`: Alexa asked the user for the value of a specific slot. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#elicitslot)
* `Dialog.ConfirmSlot`: Alexa confirmed the value of a specific slot before continuing with the dialog. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmslot)
* `Dialog.ConfirmIntent`: Alexa confirmed the all the information the user has provided for the intent before the skill took action. (https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html#confirmintent)
*/
type DialogActName string

func DialogActName_DialogElicitSlot() DialogActName {
	return "Dialog.ElicitSlot"
}
func DialogActName_DialogConfirmSlot() DialogActName {
	return "Dialog.ConfirmSlot"
}
func DialogActName_DialogConfirmIntent() DialogActName {
	return "Dialog.ConfirmIntent"
}
