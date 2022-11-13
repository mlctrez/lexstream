package isp

/*
ProductType Type of in-skill product.
*/
type ProductType string

func ProductType_SUBSCRIPTION() ProductType {
	return "SUBSCRIPTION"
}
func ProductType_ENTITLEMENT() ProductType {
	return "ENTITLEMENT"
}
func ProductType_CONSUMABLE() ProductType {
	return "CONSUMABLE"
}
