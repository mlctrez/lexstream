package typeversion

/*
SlotTypeUpdate Slot Type update description wrapper.
*/
type SlotTypeUpdate struct {
	SlotType *SlotTypeUpdateObject `json,omitempty:"slotType"`
}

/*
{
 "description": "Slot Type update description wrapper.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.slotTypeUpdateObject"
  }
 },
 "type": "object"
}
*/
