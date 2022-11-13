package version

/*
ValueSchema The value schema in type object of interaction model.
*/
type ValueSchema struct {
	Id   string           `json,omitempty:"id"`
	Name *ValueSchemaName `json,omitempty:"name"`
}

/*
{
 "description": "The value schema in type object of interaction model.",
 "properties": {
  "id": {
   "type": "string"
  },
  "name": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.ValueSchemaName"
  }
 },
 "type": "object"
}
*/
