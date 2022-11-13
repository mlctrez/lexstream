package manifest

type PermissionItems struct {
	Name *PermissionName `json:"name"`
}

/*
{
 "properties": {
  "name": {
   "$ref": "#/definitions/v1.skill.Manifest.PermissionName",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
