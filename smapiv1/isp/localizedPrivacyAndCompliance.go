package isp

/*
LocalizedPrivacyAndCompliance Defines the structure for localized privacy and compliance.
*/
type LocalizedPrivacyAndCompliance struct {
	// Link to the privacy policy that applies to this in-skill product.
	PrivacyPolicyUrl string `json,omitempty:"privacyPolicyUrl"`
}

/*
{
 "description": "Defines the structure for localized privacy and compliance.",
 "properties": {
  "privacyPolicyUrl": {
   "description": "Link to the privacy policy that applies to this in-skill product.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
