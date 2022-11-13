package auditlogs

/*
Operation Object containing name and version.
*/
type Operation struct {
	Name    string `json,omitempty:"name"`
	Version string `json,omitempty:"version"`
}

/*
{
 "description": "Object containing name and version.",
 "properties": {
  "name": {
   "type": "string"
  },
  "version": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
