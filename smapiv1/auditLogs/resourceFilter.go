package auditlogs

/*
ResourceFilter Resource that the developer operated on. Both do not need to be provided.
*/
type ResourceFilter struct {
	Id    string            `json,omitempty:"id"`
	Type_ *ResourceTypeEnum `json,omitempty:"type"`
}

/*
{
 "description": "Resource that the developer operated on. Both do not need to be provided.",
 "properties": {
  "id": {
   "format": "UUID",
   "type": "string"
  },
  "type": {
   "$ref": "#/definitions/v1.auditLogs.ResourceTypeEnum",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
