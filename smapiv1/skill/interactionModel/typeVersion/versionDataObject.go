package typeversion

/*
VersionDataObject Slot Type version fields with specific data.
*/
type VersionDataObject struct {
	Definition *ValueSupplierObject `json,omitempty:"definition"`
	// Description string for specific slot type version.
	Description string `json,omitempty:"description"`
}

/*
{
 "description": "Slot Type version fields with specific data.",
 "properties": {
  "definition": {
   "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.ValueSupplierObject"
  },
  "description": {
   "description": "Description string for specific slot type version.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
