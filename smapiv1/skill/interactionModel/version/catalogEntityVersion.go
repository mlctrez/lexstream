package version

import "time"

/*
CatalogEntityVersion Version metadata about the catalog entity version.
*/
type CatalogEntityVersion struct {
	Version      string    `json,omitempty:"version"`
	CreationTime time.Time `json,omitempty:"creationTime"`
	Description  string    `json,omitempty:"description"`
	Links        *Links    `json,omitempty:"_links"`
}

/*
{
 "description": "Version metadata about the catalog entity version.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.Links"
  },
  "creationTime": {
   "format": "date-time",
   "type": "string"
  },
  "description": {
   "type": "string"
  },
  "version": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
