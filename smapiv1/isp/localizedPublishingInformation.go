package isp

/*
LocalizedPublishingInformation Defines the structure for locale specific publishing information in the in-skill product definition.
*/
type LocalizedPublishingInformation struct {
	// Name of the in-skill product that is heard by customers and displayed in the Alexa app.
	Name string `json:"name,omitempty"`
	// Uri for the small icon image of the in-skill product.
	SmallIconUri string `json:"smallIconUri,omitempty"`
	// Uri for the large icon image of the in-skill product.
	LargeIconUri string `json:"largeIconUri,omitempty"`
	// Short description of the in-skill product that displays on the in-skill product list page in the Alexa App.
	Summary string `json:"summary,omitempty"`
	// Description of the in-skill product's purpose and features, and how it works. Should describe any prerequisites like hardware or account requirements and detailed steps for the customer to get started. This description displays to customers on the in-skill product detail card in the Alexa app.
	Description string `json:"description,omitempty"`
	// Example phrases appear on the in-skill product detail page and are the key utterances that customers can say to interact directly with the in-skill product.
	ExamplePhrases []string `json:"examplePhrases,omitempty"`
	// Search terms that can be used to describe the in-skill product. This helps customers find an in-skill product.
	Keywords             []string              `json:"keywords,omitempty"`
	CustomProductPrompts *CustomProductPrompts `json:"customProductPrompts,omitempty"`
}

/*
{
 "description": "Defines the structure for locale specific publishing information in the in-skill product definition.",
 "properties": {
  "customProductPrompts": {
   "$ref": "#/definitions/v1.isp.CustomProductPrompts"
  },
  "description": {
   "description": "Description of the in-skill product's purpose and features, and how it works. Should describe any prerequisites like hardware or account requirements and detailed steps for the customer to get started. This description displays to customers on the in-skill product detail card in the Alexa app.",
   "type": "string"
  },
  "examplePhrases": {
   "description": "Example phrases appear on the in-skill product detail page and are the key utterances that customers can say to interact directly with the in-skill product.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "keywords": {
   "description": "Search terms that can be used to describe the in-skill product. This helps customers find an in-skill product.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "largeIconUri": {
   "description": "Uri for the large icon image of the in-skill product.",
   "type": "string"
  },
  "name": {
   "description": "Name of the in-skill product that is heard by customers and displayed in the Alexa app.",
   "type": "string"
  },
  "smallIconUri": {
   "description": "Uri for the small icon image of the in-skill product.",
   "type": "string"
  },
  "summary": {
   "description": "Short description of the in-skill product that displays on the in-skill product list page in the Alexa App.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
