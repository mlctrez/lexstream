package vendormanagement

/*
Vendors List of Vendors.
*/
type Vendors struct {
	Vendors []*Vendor `json:"vendors,omitempty"`
}

/*
{
 "description": "List of Vendors.",
 "properties": {
  "vendors": {
   "items": {
    "$ref": "#/definitions/v1.vendorManagement.Vendor"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
