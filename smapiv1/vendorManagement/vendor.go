package vendormanagement

/*
Vendor Vendor Response Object.
*/
type Vendor struct {
	// Name of vendor.
	Name string `json:"name"`
	// Unique identifier of vendor associated with the API caller account.
	Id string `json:"id"`
	// Roles that user has for this vendor.
	Roles []string `json:"roles"`
}

/*
{
 "description": "Vendor Response Object.",
 "properties": {
  "id": {
   "description": "Unique identifier of vendor associated with the API caller account.",
   "type": "string"
  },
  "name": {
   "description": "Name of vendor.",
   "type": "string"
  },
  "roles": {
   "description": "Roles that user has for this vendor.",
   "items": {
    "type": "string"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/