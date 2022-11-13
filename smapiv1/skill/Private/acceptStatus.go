package private

/*
AcceptStatus Enterprise IT administrators' action on the private distribution.
*/
type AcceptStatus string

func AcceptStatus_ACCEPTED() AcceptStatus {
	return "ACCEPTED"
}
func AcceptStatus_PENDING() AcceptStatus {
	return "PENDING"
}
