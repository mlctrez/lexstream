package isp

/*
SummaryPriceListing Price listing information for in-skill product.
*/
type SummaryPriceListing struct {
	// The price of an in-skill product.
	Price int `json:"price,omitempty"`
	// The prime price of an in-skill product.
	PrimeMemberPrice int       `json:"primeMemberPrice,omitempty"`
	Currency         *Currency `json:"currency,omitempty"`
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
   "description": "The price of an in-skill product.",
   "type": "number"
  },
  "primeMemberPrice": {
   "description": "The prime price of an in-skill product.",
   "type": "number"
  }
 },
 "type": "object"
}
*/
