package type1

/*
SlotTypeResponse Slot Type information.
*/
type SlotTypeResponse struct {
	SlotType *SlotTypeResponseEntity `json:"slotType"`
}

/*
{
 "description": "Slot Type information.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeResponseEntity"
  }
 },
 "type": "object"
}
*/
