package isp

/*
PurchasableState Whether or not the in-skill product is purchasable by customers. A product that is not purchasable will prevent new customers from being prompted to purchase the product. Customers who already own the product will see no effect and continue to have access to the product features.
*/
type PurchasableState string

func PurchasableState_PURCHASABLE() PurchasableState {
	return "PURCHASABLE"
}
func PurchasableState_NOT_PURCHASABLE() PurchasableState {
	return "NOT_PURCHASABLE"
}
