package type1

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
SlotTypeItem Definition for slot type entity.
*/
type SlotTypeItem struct {
	// Name of the slot type.
	Name string `json,omitempty:"name"`
	// Description string about the slot type.
	Description string `json,omitempty:"description"`
	// Identifier of the slot type, optional in get response as the request already has slotTypeId.
	Id    string         `json,omitempty:"id"`
	Links *smapiv1.Links `json,omitempty:"_links"`
}

/*
{
 "description": "Definition for slot type entity.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "description": {
   "description": "Description string about the slot type.",
   "type": "string"
  },
  "id": {
   "description": "Identifier of the slot type, optional in get response as the request already has slotTypeId.",
   "type": "string"
  },
  "name": {
   "description": "Name of the slot type.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
