package version

import "time"

/*
CatalogEntityVersion Version metadata about the catalog entity version.
*/
type CatalogEntityVersion struct {
	Version      string    `json:"version"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	Links        *Links    `json:"_links"`
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
