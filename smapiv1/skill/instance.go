package skill

/*
Instance Structure representing properties of an instance of data. Definition will be either one of a booleanInstance, stringInstance, integerInstance, or compoundInstance.
*/
type Instance struct {
	// Path that uniquely identifies the instance in the resource.
	PropertyPath string               `json:"propertyPath"`
	DataType     *ValidationDataTypes `json:"dataType"`
	// String value of the instance. Incase of null, object or array the value field would be null or not present. Incase of string, boolean or integer dataType it will be the corresponding String value.
	Value string `json:"value"`
}

/*
{
 "description": "Structure representing properties of an instance of data. Definition will be either one of a booleanInstance, stringInstance, integerInstance, or compoundInstance.",
 "properties": {
  "dataType": {
   "$ref": "#/definitions/v1.skill.ValidationDataTypes",
   "x-isEnum": true
  },
  "propertyPath": {
   "description": "Path that uniquely identifies the instance in the resource.",
   "type": "string"
  },
  "value": {
   "description": "String value of the instance. Incase of null, object or array the value field would be null or not present. Incase of string, boolean or integer dataType it will be the corresponding String value.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
