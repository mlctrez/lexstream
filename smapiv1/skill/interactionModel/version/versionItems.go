package version

/*
VersionItems Version metadata about the entity.
*/
type VersionItems struct {
	Version      string `json:"version,omitempty"`
	CreationTime string `json:"creationTime,omitempty"`
	Description  string `json:"description,omitempty"`
	Links        *Links `json:"_links,omitempty"`
}

/*
{
 "description": "Version metadata about the entity.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.skill.interactionModel.version.Links"
  },
  "creationTime": {
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
