package interactionmodel

/*
LanguageModel Define the language model.
*/
type LanguageModel struct {
	// Invocation name of the skill.
	InvocationName     string              `json:"invocationName"`
	Types              []*SlotType         `json:"types"`
	Intents            []*Intent           `json:"intents"`
	ModelConfiguration *ModelConfiguration `json:"modelConfiguration"`
}

/*
{
 "description": "Define the language model.",
 "properties": {
  "intents": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.Intent"
   },
   "type": "array"
  },
  "invocationName": {
   "description": "Invocation name of the skill.",
   "type": "string"
  },
  "modelConfiguration": {
   "$ref": "#/definitions/v1.skill.interactionModel.ModelConfiguration"
  },
  "types": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.SlotType"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
