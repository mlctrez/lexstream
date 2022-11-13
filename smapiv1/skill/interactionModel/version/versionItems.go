package version

/*
VersionItems Version metadata about the entity.
*/
type VersionItems struct {
	Version      string `json:"version"`
	CreationTime string `json:"creationTime"`
	Description  string `json:"description"`
	Links        *Links `json:"_links"`
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