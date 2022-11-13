package interactionmodel

/*
SlotValidation Validation on a slot with support for prompt and confirmation.
*/
type SlotValidation struct {
	// The exact type of validation e.g. isLessThan,isGreaterThan etc.
	Type_ string `json,omitempty:"type"`
	// The prompt id that should be used if validation fails.
	Prompt string `json,omitempty:"prompt"`
}

/*
{
 "description": "Validation on a slot with support for prompt and confirmation.",
 "discriminator": "type",
 "properties": {
  "prompt": {
   "description": "The prompt id that should be used if validation fails.",
   "type": "string"
  },
  "type": {
   "description": "The exact type of validation e.g. isLessThan,isGreaterThan etc.",
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "required": [
  "prompt",
  "type"
 ],
 "type": "object"
}
*/
