package auditlogs

/*
Operation Object containing name and version.
*/
type Operation struct {
	Name    string `json:"name"`
	Version string `json:"version"`
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
