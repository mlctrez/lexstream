package interactionmodel

type Prompt struct {
	// The prompt id, this id can be used from dialog definitions.
	Id string `json:"id,omitempty"`
	// List of variations of the prompt, each variation can be either a text string or a well defined ssml string depending on the type defined.
	Variations []*PromptItems `json:"variations,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "The prompt id, this id can be used from dialog definitions.",
   "type": "string"
  },
  "variations": {
   "description": "List of variations of the prompt, each variation can be either a text string or a well defined ssml string depending on the type defined.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.PromptItems"
   },
   "type": "array"
  }
 },
 "required": [
  "id",
  "variations"
 ],
 "type": "object"
}
*/
