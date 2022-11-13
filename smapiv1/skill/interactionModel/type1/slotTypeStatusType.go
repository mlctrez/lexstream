package type1

/*
SlotTypeStatusType Status of last modification request for a resource.
*/
type SlotTypeStatusType string

func SlotTypeStatusType_FAILED() SlotTypeStatusType {
	return "FAILED"
}
func SlotTypeStatusType_IN_PROGRESS() SlotTypeStatusType {
	return "IN_PROGRESS"
}
func SlotTypeStatusType_SUCCEEDED() SlotTypeStatusType {
	return "SUCCEEDED"
}
