package interactionmodel

type InteractionModelData struct {
	Version          string                  `json:"version,omitempty"`
	Description      string                  `json:"description,omitempty"`
	InteractionModel *InteractionModelSchema `json:"interactionModel,omitempty"`
}

/*
{
 "properties": {
  "description": {
   "type": "string"
  },
  "interactionModel": {
   "$ref": "#/definitions/v1.skill.interactionModel.InteractionModelSchema"
  },
  "version": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
