package isp

/*
PublishingInformation Defines the structure for in-skill product publishing information.
*/
type PublishingInformation struct {
	// Defines the structure for locale specific publishing information for an in-skill product.
	Locales map[string]LocalizedPublishingInformation `json:"locales,omitempty"`
	// List of countries where the in-skill product is available.
	DistributionCountries []*DistributionCountries `json:"distributionCountries,omitempty"`
	// Defines the structure for in-skill product pricing.
	Pricing        map[string]MarketplacePricing `json:"pricing,omitempty"`
	TaxInformation *TaxInformation               `json:"taxInformation,omitempty"`
}

/*
{
 "description": "Defines the structure for in-skill product publishing information.",
 "properties": {
  "distributionCountries": {
   "description": "List of countries where the in-skill product is available.",
   "items": {
    "$ref": "#/definitions/v1.isp.DistributionCountries"
   },
   "type": "array"
  },
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.isp.LocalizedPublishingInformation"
   },
   "description": "Defines the structure for locale specific publishing information for an in-skill product.",
   "type": "object"
  },
  "pricing": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.isp.MarketplacePricing"
   },
   "description": "Defines the structure for in-skill product pricing.",
   "type": "object"
  },
  "taxInformation": {
   "$ref": "#/definitions/v1.isp.TaxInformation"
  }
 },
 "type": "object"
}
*/
