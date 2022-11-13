package isp

/*
TaxInformation Defines the structure for in-skill product tax information.
*/
type TaxInformation struct {
	Category *TaxInformationCategory `json:"category,omitempty"`
}

/*
{
 "description": "Defines the structure for in-skill product tax information.",
 "properties": {
  "category": {
   "$ref": "#/definitions/v1.isp.TaxInformationCategory",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
