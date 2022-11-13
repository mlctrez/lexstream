package interactionmodel

type InteractionModelSchema struct {
	LanguageModel *LanguageModel `json:"languageModel"`
	Dialog        *Dialog        `json:"dialog"`
	// List of prompts.
	Prompts []*Prompt `json:"prompts"`
}

/*
{
 "properties": {
  "dialog": {
   "$ref": "#/definitions/v1.skill.interactionModel.Dialog"
  },
  "languageModel": {
   "$ref": "#/definitions/v1.skill.interactionModel.languageModel"
  },
  "prompts": {
   "description": "List of prompts.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.Prompt"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/