package type1

/*
SlotTypeStatus Defines the structure for slot type status response.
*/
type SlotTypeStatus struct {
	UpdateRequest *LastUpdateRequest `json:"updateRequest"`
}

/*
{
 "description": "Defines the structure for slot type status response.",
 "properties": {
  "updateRequest": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.LastUpdateRequest"
  }
 },
 "type": "object"
}
*/
