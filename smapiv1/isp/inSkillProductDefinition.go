package isp

/*
InSkillProductDefinition Defines the structure for an in-skill product.
*/
type InSkillProductDefinition struct {
	// Version of in-skill product definition.
	Version string       `json:"version,omitempty"`
	Type_   *ProductType `json:"type,omitempty"`
	// Developer selected in-skill product name. This is for developer reference only, it can be used to filter query results to identify a matching in-skill product.
	ReferenceName           string                   `json:"referenceName,omitempty"`
	PurchasableState        *PurchasableState        `json:"purchasableState,omitempty"`
	PromotableState         *PromotableState         `json:"promotableState,omitempty"`
	SubscriptionInformation *SubscriptionInformation `json:"subscriptionInformation,omitempty"`
	PublishingInformation   *PublishingInformation   `json:"publishingInformation,omitempty"`
	PrivacyAndCompliance    *PrivacyAndCompliance    `json:"privacyAndCompliance,omitempty"`
	// Special instructions provided by the developer to test the in-skill product.
	TestingInstructions string `json:"testingInstructions,omitempty"`
}

/*
{
 "description": "Defines the structure for an in-skill product.",
 "properties": {
  "privacyAndCompliance": {
   "$ref": "#/definitions/v1.isp.PrivacyAndCompliance"
  },
  "promotableState": {
   "$ref": "#/definitions/v1.isp.PromotableState",
   "x-isEnum": true
  },
  "publishingInformation": {
   "$ref": "#/definitions/v1.isp.PublishingInformation"
  },
  "purchasableState": {
   "$ref": "#/definitions/v1.isp.PurchasableState",
   "x-isEnum": true
  },
  "referenceName": {
   "description": "Developer selected in-skill product name. This is for developer reference only, it can be used to filter query results to identify a matching in-skill product.",
   "type": "string"
  },
  "subscriptionInformation": {
   "$ref": "#/definitions/v1.isp.SubscriptionInformation"
  },
  "testingInstructions": {
   "description": "Special instructions provided by the developer to test the in-skill product.",
   "type": "string"
  },
  "type": {
   "$ref": "#/definitions/v1.isp.ProductType",
   "x-isEnum": true
  },
  "version": {
   "description": "Version of in-skill product definition.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
