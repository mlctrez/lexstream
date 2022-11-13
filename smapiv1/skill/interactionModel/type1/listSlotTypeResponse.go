package type1

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListSlotTypeResponse List of slot types of a skill for the vendor.
*/
type ListSlotTypeResponse struct {
	Links *smapiv1.Links `json:"_links"`
	/*
	   List of slot types.
	*/
	SlotTypes []*SlotTypeItem `json:"slotTypes"`
	NextToken string          `json:"nextToken"`
}

/*
{
 "description": "List of slot types of a skill for the vendor.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "nextToken": {
   "type": "string"
  },
  "slotTypes": {
   "description": "List of slot types.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.type.SlotTypeItem"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
