package type1

/*
DefinitionData Slot type request definitions.
*/
type DefinitionData struct {
	SlotType *SlotTypeInput `json,omitempty:"slotType"`
	// The vendorId that the slot type should belong to.
	VendorId string `json,omitempty:"vendorId"`
}

/*
{
 "description": "Slot type request definitions.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeInput"
  },
  "vendorId": {
   "description": "The vendorId that the slot type should belong to.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
