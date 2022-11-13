package version

/*
VersionItems Version metadata about the entity.
*/
type VersionItems struct {
	Version      string `json,omitempty:"version"`
	CreationTime string `json,omitempty:"creationTime"`
	Description  string `json,omitempty:"description"`
	Links        *Links `json,omitempty:"_links"`
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
