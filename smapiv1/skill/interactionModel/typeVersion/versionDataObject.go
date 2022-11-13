package typeversion

/*
VersionDataObject Slot Type version fields with specific data.
*/
type VersionDataObject struct {
	Definition *ValueSupplierObject `json:"definition,omitempty"`
	// Description string for specific slot type version.
	Description string `json:"description,omitempty"`
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
