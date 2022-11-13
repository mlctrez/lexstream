package interactionmodel

/*
TypeValue The value schema in type object of interaction model.
*/
type TypeValue struct {
	Id   string           `json:"id"`
	Name *TypeValueObject `json:"name"`
}

/*
{
 "description": "The value schema in type object of interaction model.",
 "properties": {
  "id": {
   "type": "string"
  },
  "name": {
   "$ref": "#/definitions/v1.skill.interactionModel.TypeValueObject"
  }
 },
 "type": "object"
}
*/
