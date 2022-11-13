package type1

/*
Error The error which would fail requests.
*/
type Error struct {
	// The error code.
	Code string `json,omitempty:"code"`
	// The error message.
	Message string `json,omitempty:"message"`
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
