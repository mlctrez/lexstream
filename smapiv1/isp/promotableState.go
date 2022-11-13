package isp

/*
PromotableState Promote this ISP on Amazon channels such as Amazon.com. Enabling this setting will allow customers to view ISP detail pages and purchase the ISP on Amazon.com.
*/
type PromotableState string

func PromotableState_IN_SKILL_ONLY() PromotableState {
	return "IN_SKILL_ONLY"
}
func PromotableState_ALL_AMAZON_CHANNELS() PromotableState {
	return "ALL_AMAZON_CHANNELS"
}
