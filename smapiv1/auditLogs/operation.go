package auditlogs

/*
Operation Object containing name and version.
*/
type Operation struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
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
