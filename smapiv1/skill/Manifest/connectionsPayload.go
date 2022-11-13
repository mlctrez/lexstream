package manifest

/*
ConnectionsPayload Payload of the connection.
*/
type ConnectionsPayload struct {
	// Type of the payload.
	Type_ string `json:"type"`
	// Version of the payload.
	Version string `json:"version"`
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
