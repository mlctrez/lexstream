package type1

/*
LastUpdateRequest Contains attributes related to last modification request of a resource.
*/
type LastUpdateRequest struct {
	Status *SlotTypeStatusType `json:"status"`
	// The version id of the entity returned.
	Version  string     `json:"version"`
	Errors   []*Error   `json:"errors"`
	Warnings []*Warning `json:"warnings"`
}

/*
{
 "description": "Contains attributes related to last modification request of a resource.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.Error"
   },
   "type": "array"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeStatusType",
   "x-isEnum": true
  },
  "version": {
   "description": "The version id of the entity returned.",
   "type": "string"
  },
  "warnings": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.Warning"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
