package typeversion

/*
SlotTypeVersionDataObject Slot Type version fields with metadata.
*/
type SlotTypeVersionDataObject struct {
	// Slot type id associated with the slot type version.
	Id         string               `json:"id,omitempty"`
	Definition *ValueSupplierObject `json:"definition,omitempty"`
	// Description string for specific slot type version.
	Description string `json:"description,omitempty"`
	// Specific slot type version.
	Version string `json:"version,omitempty"`
}

/*
{
 "description": "Slot Type version fields with metadata.",
 "properties": {
  "definition": {
   "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.ValueSupplierObject"
  },
  "description": {
   "description": "Description string for specific slot type version.",
   "type": "string"
  },
  "id": {
   "description": "Slot type id associated with the slot type version.",
   "type": "string"
  },
  "version": {
   "description": "Specific slot type version.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
