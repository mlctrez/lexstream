package isp

/*
SubscriptionPaymentFrequency The frequency in which payments are collected for the subscription.
*/
type SubscriptionPaymentFrequency string

func SubscriptionPaymentFrequency_MONTHLY() SubscriptionPaymentFrequency {
	return "MONTHLY"
}
func SubscriptionPaymentFrequency_YEARLY() SubscriptionPaymentFrequency {
	return "YEARLY"
}
