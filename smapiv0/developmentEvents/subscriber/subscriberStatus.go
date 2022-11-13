package subscriber

/*
SubscriberStatus Status of the subscriber. This enum may get extended with new values in future. Clients are expected to gracefully handle any unknown values.
*/
type SubscriberStatus string

func SubscriberStatus_ACTIVE() SubscriberStatus {
	return "ACTIVE"
}
func SubscriberStatus_INACTIVE() SubscriberStatus {
	return "INACTIVE"
}
