package subscriber

type CreateSubscriberRequest struct {
	// Name of the subscriber.
	Name string `json:"name"`
	// The Vendor ID.
	VendorId string    `json:"vendorId"`
	Endpoint *Endpoint `json:"endpoint"`
}

/*
{
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.Endpoint"
  },
  "name": {
   "description": "Name of the subscriber.",
   "type": "string"
  },
  "vendorId": {
   "description": "The Vendor ID.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
