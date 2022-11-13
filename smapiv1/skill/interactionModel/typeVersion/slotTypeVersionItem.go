package typeversion

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
SlotTypeVersionItem Definition for slot type entity.
*/
type SlotTypeVersionItem struct {
	// Version number of slot type.
	Version string `json:"version"`
	// Description string about the slot type version.
	Description string         `json:"description"`
	Links       *smapiv1.Links `json:"_links"`
}

/*
{
 "description": "Definition for slot type entity.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "description": {
   "description": "Description string about the slot type version.",
   "type": "string"
  },
  "version": {
   "description": "Version number of slot type.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
