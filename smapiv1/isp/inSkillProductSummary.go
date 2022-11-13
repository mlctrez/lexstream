package isp

import "time"

/*
InSkillProductSummary Information about the in-skill product that is not editable.
*/
type InSkillProductSummary struct {
	Type_ *ProductType `json:"type"`
	// primary identifier of in-skill product.
	ProductId string `json:"productId"`
	// Developer selected in-skill product name. This is for developer reference only, it can be used to filter query results to identify a matching in-skill product.
	ReferenceName string `json:"referenceName"`
	// Date of last update.
	LastUpdated      time.Time         `json:"lastUpdated"`
	NameByLocale     map[string]string `json:"nameByLocale"`
	Status           *Status           `json:"status"`
	Stage            *Stage            `json:"stage"`
	EditableState    *EditableState    `json:"editableState"`
	PurchasableState *PurchasableState `json:"purchasableState"`
	PromotableState  *PromotableState  `json:"promotableState"`
	Links            *IspSummaryLinks  `json:"_links"`
	// In-skill product pricing information.
	Pricing map[string]SummaryMarketplacePricing `json:"pricing"`
}

/*
{
 "description": "Information about the in-skill product that is not editable.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.isp.IspSummaryLinks"
  },
  "editableState": {
   "$ref": "#/definitions/v1.isp.EditableState",
   "x-isEnum": true
  },
  "lastUpdated": {
   "description": "Date of last update.",
   "format": "date-time",
   "type": "string"
  },
  "nameByLocale": {
   "additionalProperties": {
    "type": "string"
   },
   "type": "object"
  },
  "pricing": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.isp.SummaryMarketplacePricing"
   },
   "description": "In-skill product pricing information.",
   "type": "object"
  },
  "productId": {
   "description": "primary identifier of in-skill product.",
   "type": "string"
  },
  "promotableState": {
   "$ref": "#/definitions/v1.isp.PromotableState",
   "x-isEnum": true
  },
  "purchasableState": {
   "$ref": "#/definitions/v1.isp.PurchasableState",
   "x-isEnum": true
  },
  "referenceName": {
   "description": "Developer selected in-skill product name. This is for developer reference only, it can be used to filter query results to identify a matching in-skill product.",
   "type": "string"
  },
  "stage": {
   "$ref": "#/definitions/v1.isp.Stage",
   "x-isEnum": true
  },
  "status": {
   "$ref": "#/definitions/v1.isp.Status",
   "x-isEnum": true
  },
  "type": {
   "$ref": "#/definitions/v1.isp.ProductType",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
