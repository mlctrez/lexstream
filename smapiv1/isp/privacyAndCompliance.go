package isp

/*
PrivacyAndCompliance Defines the structure for privacy and compliance.
*/
type PrivacyAndCompliance struct {
	// Defines the structure for locale specific privacy and compliance.
	Locales map[string]LocalizedPrivacyAndCompliance `json:"locales"`
}

/*
{
 "description": "Defines the structure for privacy and compliance.",
 "properties": {
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.isp.LocalizedPrivacyAndCompliance"
   },
   "description": "Defines the structure for locale specific privacy and compliance.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
