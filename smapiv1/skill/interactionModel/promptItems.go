package interactionmodel

type PromptItems struct {
	Type_ *PromptItemsType `json:"type"`
	// Specifies the prompt.
	Value string `json:"value"`
}

/*
{
 "properties": {
  "type": {
   "$ref": "#/definitions/v1.skill.interactionModel.PromptItemsType",
   "x-isEnum": true
  },
  "value": {
   "description": "Specifies the prompt.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
