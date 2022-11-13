package type1

/*
SlotTypeDefinitionOutput Slot Type request definitions.
*/
type SlotTypeDefinitionOutput struct {
	SlotType *SlotTypeInput `json,omitempty:"slotType"`
	// Total number of versions.
	TotalVersions string `json,omitempty:"totalVersions"`
}

/*
{
 "description": "Slot Type request definitions.",
 "properties": {
  "slotType": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeInput"
  },
  "totalVersions": {
   "description": "Total number of versions.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
