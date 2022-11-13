package manifest

/*
ConnectionsPayload Payload of the connection.
*/
type ConnectionsPayload struct {
	// Type of the payload.
	Type_ string `json,omitempty:"type"`
	// Version of the payload.
	Version string `json,omitempty:"version"`
}

/*
{
 "description": "Payload of the connection.",
 "properties": {
  "type": {
   "description": "Type of the payload.",
   "type": "string"
  },
  "version": {
   "description": "Version of the payload.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
