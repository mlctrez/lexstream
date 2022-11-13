package isp

/*
Currency Currency to use for in-skill product.
*/
type Currency string

func Currency_USD() Currency {
	return "USD"
}
func Currency_GBP() Currency {
	return "GBP"
}
func Currency_EUR() Currency {
	return "EUR"
}
func Currency_JPY() Currency {
	return "JPY"
}
