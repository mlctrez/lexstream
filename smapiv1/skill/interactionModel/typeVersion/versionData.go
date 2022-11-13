package typeversion

/*
VersionData Slot Type version specific data.
*/
type VersionData struct {
	SlotType *VersionDataObject `json:"slotType"`
}

/*
{
 "description": "Slot Type version specific data.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.VersionDataObject"
  }
 },
 "type": "object"
}
*/
