package typeversion

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListSlotTypeVersionResponse List of slot type versions of a skill for the vendor.
*/
type ListSlotTypeVersionResponse struct {
	Links *smapiv1.Links `json:"_links"`
	/*
	   List of slot types.
	*/
	SlotTypeVersions []*SlotTypeVersionItem `json:"slotTypeVersions"`
	NextToken        string                 `json:"nextToken"`
}

/*
{
 "description": "List of slot type versions of a skill for the vendor.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "nextToken": {
   "type": "string"
  },
  "slotTypeVersions": {
   "description": "List of slot types.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.typeVersion.SlotTypeVersionItem"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
