package skill

type CreateSkillWithPackageRequest struct {
	// ID of the vendor owning the skill.
	VendorId string `json:"vendorId"`
	// The URL for the skill package.
	Location string `json:"location"`
}

/*
{
 "properties": {
  "location": {
   "description": "The URL for the skill package.",
   "format": "uri",
   "type": "string"
  },
  "vendorId": {
   "description": "ID of the vendor owning the skill.",
   "type": "string"
  }
 },
 "required": [
  "location"
 ],
 "type": "object"
}
*/
