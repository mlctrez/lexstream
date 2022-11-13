package interactionmodel

type InteractionModelData struct {
	Version          string                  `json:"version"`
	Description      string                  `json:"description"`
	InteractionModel *InteractionModelSchema `json:"interactionModel"`
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
