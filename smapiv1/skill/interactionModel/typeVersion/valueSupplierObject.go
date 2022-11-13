package typeversion

import interactionmodel "github.com/mlctrez/lexstream/smapiv1/skill/interactionModel"

/*
ValueSupplierObject Value supplier object for slot definition.
*/
type ValueSupplierObject struct {
	ValueSupplier *interactionmodel.ValueSupplier `json,omitempty:"valueSupplier"`
}

/*
{
 "description": "Value supplier object for slot definition.",
 "properties": {
  "valueSupplier": {
   "$ref": "#/definitions/v1.skill.interactionModel.ValueSupplier"
  }
 },
 "type": "object"
}
*/
