package isp

import "time"

/*
SummaryMarketplacePricing Localized in-skill product pricing information.
*/
type SummaryMarketplacePricing struct {
	// Date when in-skill product is available to customers for both purchase and use. Prior to this date the in-skill product will appear unavailable to customers and will not be purchasable.
	ReleaseDate         time.Time            `json,omitempty:"releaseDate"`
	DefaultPriceListing *SummaryPriceListing `json,omitempty:"defaultPriceListing"`
}

/*
{
 "description": "Localized in-skill product pricing information.",
 "properties": {
  "defaultPriceListing": {
   "$ref": "#/definitions/v1.isp.SummaryPriceListing"
  },
  "releaseDate": {
   "description": "Date when in-skill product is available to customers for both purchase and use. Prior to this date the in-skill product will appear unavailable to customers and will not be purchasable.",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
