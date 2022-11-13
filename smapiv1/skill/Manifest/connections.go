package manifest

/*
Connections Skill connection object.
*/
type Connections struct {
	// Name of the connection.
	Name    string              `json,omitempty:"name"`
	Payload *ConnectionsPayload `json,omitempty:"payload"`
}

/*
{
 "description": "Skill connection object.",
 "properties": {
  "name": {
   "description": "Name of the connection.",
   "type": "string"
  },
  "payload": {
   "$ref": "#/definitions/v1.skill.Manifest.ConnectionsPayload"
  }
 },
 "type": "object"
}
*/
