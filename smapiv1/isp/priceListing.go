package isp

/*
PriceListing Price listing information for in-skill product.
*/
type PriceListing struct {
	// Defines the price of an in-skill product. The list price should be your suggested price, not including any VAT or similar taxes. Taxes are included in the final price to end users.
	Price    int       `json:"price"`
	Currency *Currency `json:"currency"`
}

/*
{
 "description": "Price listing information for in-skill product.",
 "properties": {
  "currency": {
   "$ref": "#/definitions/v1.isp.Currency",
   "x-isEnum": true
  },
  "price": {
   "description": "Defines the price of an in-skill product. The list price should be your suggested price, not including any VAT or similar taxes. Taxes are included in the final price to end users.",
   "type": "number"
  }
 },
 "type": "object"
}
*/
