package subscriber

type CreateSubscriberRequest struct {
	// Name of the subscriber.
	Name string `json:"name,omitempty"`
	// The Vendor ID.
	VendorId string    `json:"vendorId,omitempty"`
	Endpoint *Endpoint `json:"endpoint,omitempty"`
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
