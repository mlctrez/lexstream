package isp

type CreateInSkillProductRequest struct {
	// ID of the vendor owning the in-skill product.
	VendorId                 string                    `json:"vendorId,omitempty"`
	InSkillProductDefinition *InSkillProductDefinition `json:"inSkillProductDefinition,omitempty"`
}

/*
{
 "properties": {
  "inSkillProductDefinition": {
   "$ref": "#/definitions/v1.isp.InSkillProductDefinition"
  },
  "vendorId": {
   "description": "ID of the vendor owning the in-skill product.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
