package isp

import "time"

/*
MarketplacePricing In-skill product pricing information for a marketplace.
*/
type MarketplacePricing struct {
	// Date when in-skill product is available to customers for both purchase and use. Prior to this date the in-skill product will appear unavailable to customers and will not be purchasable.
	ReleaseDate         time.Time     `json,omitempty:"releaseDate"`
	DefaultPriceListing *PriceListing `json,omitempty:"defaultPriceListing"`
}

/*
{
 "description": "In-skill product pricing information for a marketplace.",
 "properties": {
  "defaultPriceListing": {
   "$ref": "#/definitions/v1.isp.PriceListing"
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
