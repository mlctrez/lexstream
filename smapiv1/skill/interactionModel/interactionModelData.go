package interactionmodel

type InteractionModelData struct {
	Version          string                  `json,omitempty:"version"`
	Description      string                  `json,omitempty:"description"`
	InteractionModel *InteractionModelSchema `json,omitempty:"interactionModel"`
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
