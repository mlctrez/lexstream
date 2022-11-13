package isp

/*
InSkillProductDefinition Defines the structure for an in-skill product.
*/
type InSkillProductDefinition struct {
	// Version of in-skill product definition.
	Version string       `json:"version"`
	Type_   *ProductType `json:"type"`
	// Developer selected in-skill product name. This is for developer reference only, it can be used to filter query results to identify a matching in-skill product.
	ReferenceName           string                   `json:"referenceName"`
	PurchasableState        *PurchasableState        `json:"purchasableState"`
	PromotableState         *PromotableState         `json:"promotableState"`
	SubscriptionInformation *SubscriptionInformation `json:"subscriptionInformation"`
	PublishingInformation   *PublishingInformation   `json:"publishingInformation"`
	PrivacyAndCompliance    *PrivacyAndCompliance    `json:"privacyAndCompliance"`
	// Special instructions provided by the developer to test the in-skill product.
	TestingInstructions string `json:"testingInstructions"`
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
