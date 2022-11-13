package type1

/*
UpdateRequest Slot type update request object.
*/
type UpdateRequest struct {
	SlotType *SlotTypeUpdateDefinition `json,omitempty:"slotType"`
}

/*
{
 "description": "Slot type update request object.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeUpdateDefinition"
  }
 },
 "type": "object"
}
*/
