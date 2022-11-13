package type1

/*
Error The error which would fail requests.
*/
type Error struct {
	// The error code.
	Code string `json:"code,omitempty"`
	// The error message.
	Message string `json:"message,omitempty"`
}

/*
{
 "description": "The error which would fail requests.",
 "properties": {
  "code": {
   "description": "The error code.",
   "type": "string"
  },
  "message": {
   "description": "The error message.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
