package isp

/*
EditableState Whether or not the in-skill product is editable.
*/
type EditableState string

func EditableState_EDITABLE() EditableState {
	return "EDITABLE"
}
func EditableState_NOT_EDITABLE() EditableState {
	return "NOT_EDITABLE"
}
