package isp

/*
ProductResponse Product ID information.
*/
type ProductResponse struct {
	// ID of the in-skill product created.
	ProductId string `json,omitempty:"productId"`
}

/*
{
 "description": "Product ID information.",
 "properties": {
  "productId": {
   "description": "ID of the in-skill product created.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
