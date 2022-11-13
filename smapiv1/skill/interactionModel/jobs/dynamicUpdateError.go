package jobs

/*
DynamicUpdateError Error schema for dynamic update.
*/
type DynamicUpdateError struct {
	// Dynamic update error code.
	Code string `json:"code,omitempty"`
	// Readable description of error.
	Message string `json:"message,omitempty"`
}

/*
{
 "description": "Error schema for dynamic update.",
 "properties": {
  "code": {
   "description": "Dynamic update error code.",
   "type": "string"
  },
  "message": {
   "description": "Readable description of error.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
