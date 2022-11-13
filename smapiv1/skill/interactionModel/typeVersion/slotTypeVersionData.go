package typeversion

/*
SlotTypeVersionData Slot Type version data with metadata.
*/
type SlotTypeVersionData struct {
	SlotType *SlotTypeVersionDataObject `json:"slotType,omitempty"`
}

/*
{
 "description": "Slot Type version data with metadata.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.SlotTypeVersionDataObject"
  }
 },
 "type": "object"
}
*/
