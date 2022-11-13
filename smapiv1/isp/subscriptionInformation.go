package isp

/*
SubscriptionInformation Defines the structure for in-skill product subscription information.
*/
type SubscriptionInformation struct {
	SubscriptionPaymentFrequency *SubscriptionPaymentFrequency `json:"subscriptionPaymentFrequency"`
	// Days of free trial period for subscription. Max allowed is 365 days.
	SubscriptionTrialPeriodDays int `json:"subscriptionTrialPeriodDays"`
}

/*
{
 "description": "Defines the structure for in-skill product subscription information.",
 "properties": {
  "subscriptionPaymentFrequency": {
   "$ref": "#/definitions/v1.isp.subscriptionPaymentFrequency",
   "x-isEnum": true
  },
  "subscriptionTrialPeriodDays": {
   "description": "Days of free trial period for subscription. Max allowed is 365 days.",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
