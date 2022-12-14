package manifest

/*
Connections Skill connection object.
*/
type Connections struct {
	// Name of the connection.
	Name    string              `json:"name,omitempty"`
	Payload *ConnectionsPayload `json:"payload,omitempty"`
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
