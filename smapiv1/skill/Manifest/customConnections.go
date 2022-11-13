package manifest

/*
CustomConnections Supported connections.
*/
type CustomConnections struct {
	// List of required connections.
	Requires []*Connections `json,omitempty:"requires"`
	// List of provided connections.
	Provides []*Connections `json,omitempty:"provides"`
}

/*
{
 "description": "Supported connections.",
 "properties": {
  "provides": {
   "description": "List of provided connections.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.Connections"
   },
   "type": "array"
  },
  "requires": {
   "description": "List of required connections.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.Connections"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
