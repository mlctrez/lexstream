package isp

/*
CustomProductPrompts Custom prompts used for in-skill product purchasing options. Supports Speech Synthesis Markup Language (SSML), which can be used to control pronunciation, intonation, timing, and emotion.
*/
type CustomProductPrompts struct {
	// Description of in-skill product heard before customer is prompted for purchase.
	PurchasePromptDescription string `json:"purchasePromptDescription,omitempty"`
	// A description of the product that displays on the skill card in the Alexa app.
	BoughtCardDescription string `json:"boughtCardDescription,omitempty"`
}

/*
{
 "description": "Custom prompts used for in-skill product purchasing options. Supports Speech Synthesis Markup Language (SSML), which can be used to control pronunciation, intonation, timing, and emotion.",
 "properties": {
  "boughtCardDescription": {
   "description": "A description of the product that displays on the skill card in the Alexa app.",
   "example": "Enjoy {PREMIUM_CONTENT_TITLE}! Ask for a list of adventures to explore your purchase..",
   "type": "string"
  },
  "purchasePromptDescription": {
   "description": "Description of in-skill product heard before customer is prompted for purchase.",
   "example": "{PREMIUM_CONTENT_TITLE} includes \u003cemphasis level=\"moderate\"\u003e an assortment of fifty questions on a broad range of historical topics.\u003cbreak time=\"150ms\"\u003e",
   "type": "string"
  }
 },
 "type": "object"
}
*/
