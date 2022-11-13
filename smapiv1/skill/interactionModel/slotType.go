package interactionmodel

/*
SlotType Custom slot type to define a list of possible values for a slot. Used for items that are not covered by Amazon's built-in slot types.
*/
type SlotType struct {
	// The name of the custom slot type.
	Name string `json:"name"`
	// List of expected values. Values outside the list are still returned.
	Values        []*TypeValue   `json:"values"`
	ValueSupplier *ValueSupplier `json:"valueSupplier"`
}

/*
{
 "description": "Custom slot type to define a list of possible values for a slot. Used for items that are not covered by Amazon's built-in slot types.",
 "properties": {
  "name": {
   "description": "The name of the custom slot type.",
   "type": "string"
  },
  "valueSupplier": {
   "$ref": "#/definitions/v1.skill.interactionModel.ValueSupplier"
  },
  "values": {
   "description": "List of expected values. Values outside the list are still returned.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.TypeValue"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
