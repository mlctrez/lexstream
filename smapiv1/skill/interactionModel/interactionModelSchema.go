package interactionmodel

type InteractionModelSchema struct {
	LanguageModel *LanguageModel `json:"languageModel,omitempty"`
	Dialog        *Dialog        `json:"dialog,omitempty"`
	// List of prompts.
	Prompts []*Prompt `json:"prompts,omitempty"`
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
